package database_test

import (
	"context"

	"easyslip.cc/mic-project-layout/internal/database"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MockRedis", func() {
	var cleanup func()
	var cache *database.Rdb
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		cache, cleanup = database.MockRdb()
		Expect(cache).ShouldNot(BeNil())
		Expect(cleanup).ShouldNot(BeNil())

		err := cache.Clt.Ping(ctx).Err()
		Expect(err).ShouldNot(HaveOccurred())

	})

	AfterEach(func() { cleanup() })

	DescribeTable("save and get int",
		func(key string, n int) {
			err := cache.SaveInt(ctx, key, n, 0)
			Expect(err).ShouldNot(HaveOccurred())

			v, err := cache.GetInt(ctx, key)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v).Should(Equal(n))

		},
		Entry("save 1", "x", 5),
		Entry("save 2", "y", 8),
	)

	DescribeTable("save and get string",
		func(key string, s string) {
			err := cache.SaveString(ctx, key, s, 0)
			Expect(err).ShouldNot(HaveOccurred())

			v, err := cache.MustGetString(ctx, key)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v).Should(Equal(s))

		},

		Entry("save x", "a", "x"),
		Entry("save y", "b", "y"),
	)

	DescribeTable("expire save",
		func(key string, s string) {
			err := cache.SaveString(ctx, key, s, 10)
			Expect(err).ShouldNot(HaveOccurred())

			err = cache.Expire(ctx, key, 0)
			Expect(err).ShouldNot(HaveOccurred())

			v, err := cache.MustGetString(ctx, key)
			Expect(err).Should(HaveOccurred())
			Expect(errors.Cause(err)).Should(BeEquivalentTo(redis.Nil))
			Expect(v).Should(BeEmpty())

		},
		Entry("expire save x", "a1", "x"),
		Entry("expire save y", "b1", "y"),
	)

	DescribeTable("del",
		func(key string, s string) {
			err := cache.SaveString(ctx, key, s, 60)
			Expect(err).ShouldNot(HaveOccurred())

			err = cache.DelKey(ctx, key)
			Expect(err).ShouldNot(HaveOccurred())

			v, err := cache.MustGetString(ctx, key)
			Expect(err).Should(HaveOccurred())
			Expect(errors.Cause(err)).Should(BeEquivalentTo(redis.Nil))
			Expect(v).Should(BeEmpty())
		},
		Entry("del save x", "a2", "x"),
		Entry("del save y", "b2", "y"),
	)

	DescribeTable("autoCacheString",
		func(key string, s string, use bool, eq string) {
			f := func() (string, int, error) {
				return s, 120, nil
			}
			v, err := cache.AutoCacheString(ctx, key, use, f)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v).Should(Equal(s))
		},

		Entry("autoCacheString x", "key", "x", true, ""),
		Entry("autoCacheString x", "key", "x", false, "x"),
		Entry("autoCacheString x", "key", "x", true, "x"),
	)

	It("Keys", func() {
		_ = cache.SaveString(ctx, "x:a", "a", 0)
		_ = cache.SaveString(ctx, "x:b", "b", 0)
		_ = cache.SaveInt(ctx, "x:c", 1, 0)

		keys, err := cache.Keys(ctx, "x:*")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(keys).Should(HaveLen(3))
	})

})
