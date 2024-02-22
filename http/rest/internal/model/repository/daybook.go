package repository

import (
	"context"
	mDaybook "saved/http/rest/internal/model/daybook"
)

type DaybookRepository interface {
	Count(ctx context.Context, query map[string][]string) (int64, error)
	FindAll(ctx context.Context, query map[string][]string) ([]mDaybook.DaybookList, error)
	FindById(ctx context.Context, id string) (mDaybook.DaybookResponse, error)
	FindByIdForExcel(ctx context.Context, id string) (mDaybook.DaybookExpand, error)
	Create(ctx context.Context, payload mDaybook.Daybook) (mDaybook.Daybook, error)
	BulkCreateDaybookDetail(ctx context.Context, payloads []interface{}) error
	Update(ctx context.Context, payload mDaybook.Daybook) (mDaybook.Daybook, error)
}
