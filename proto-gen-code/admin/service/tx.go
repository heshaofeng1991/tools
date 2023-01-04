package service

import (
	"gorm.io/gorm"
)

func (s *Service) Tx(begin ...bool) *gorm.DB {
	isBegin := false
	if len(begin) > 0 {
		isBegin = begin[0]
	}

	if s.tx != nil {
		return s.tx
	}

	if isBegin {
		s.tx = s.db.Begin()
		return s.tx
	}

	return s.db
}

func (s *Service) SetTx(tx *gorm.DB) {
	s.tx = tx
}

func (s *Service) TxCommit() {
	s.tx.Commit()

	s.tx = nil
}
