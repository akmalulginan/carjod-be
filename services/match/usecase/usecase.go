package usecase

import (
	"context"
	"errors"

	"github.com/akmalulginan/carjod-be/domain"
)

type matchUsecase struct {
	txCoordinator   domain.TxCoordinator
	matchRepository domain.MatchRepository
	userRepository  domain.UserRepository
}

func NewMatchUsecase(mr domain.MatchRepository, ur domain.UserRepository, txC domain.TxCoordinator) domain.MatchUsecase {
	return matchUsecase{userRepository: ur, matchRepository: mr, txCoordinator: txC}
}

func (u matchUsecase) GetCandidate(ctx context.Context, userId string) (candidate domain.User, err error) {
	user, err := u.userRepository.FindById(ctx, userId)
	if err != nil {
		return user, err
	}

	matches, err := u.matchRepository.FindByUserId(ctx, userId, false)
	if err != nil {
		return user, err
	}

	matchTargetIds := make([]string, 0)
	for _, v := range matches {
		matchTargetIds = append(matchTargetIds, v.TargetUserId)
	}

	candidate, err = u.userRepository.FindCandidate(ctx, user, matchTargetIds)
	if err != nil {
		return user, err
	}

	return candidate, nil
}

func (u matchUsecase) GetMatches(ctx context.Context, userId string) (users []domain.User, err error) {
	matches, err := u.matchRepository.FindByUserId(ctx, userId, false)
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0)
	for _, v := range matches {
		ids = append(ids, v.UserId)
	}

	users, err = u.userRepository.FindByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u matchUsecase) Action(ctx context.Context, data *domain.Match) (err error) {
	user, err := u.userRepository.FindById(ctx, data.UserId)
	if err != nil {
		return err
	}

	matches, err := u.matchRepository.FindByUserId(ctx, data.UserId, true)
	if err != nil {
		return err
	}

	if len(matches) >= 10 {
		err = errors.New("daily limit has been reached")

		if !user.PremiumSwipe {
			return err
		}

		if user.PremiumSwipe && !user.PremiumActicve {
			return err
		}
	}

	ctx, err = u.txCoordinator.Begin(ctx)
	if err != nil {
		return err
	}

	defer u.txCoordinator.Rollback(ctx)

	if data.IsLike {
		match, err := u.matchRepository.FindLiked(ctx, data.TargetUserId, data.UserId)
		if err != nil {
			return err
		}

		if match.IsLike {
			match.IsMatch = true
			data.IsMatch = true

			err = u.matchRepository.Update(ctx, &match)
			if err != nil {
				return err
			}
		}

	}

	err = u.matchRepository.Create(ctx, data)
	if err != nil {
		return err
	}

	return u.txCoordinator.Commit(ctx)
}
