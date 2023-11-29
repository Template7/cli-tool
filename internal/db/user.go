package db

import "context"

func (c *Client) ListFakeUsers(ctx context.Context) (data []string) {
	log := c.log.WithContext(ctx)
	log.Debug("list fake users")

	if err := c.sql.core.WithContext(ctx).Raw("select id from template7.user where username like 'fakeuser%'").Scan(&data).Error; err != nil {
		c.log.WithError(err).Error("fail to get fake users")
	}
	return
}
