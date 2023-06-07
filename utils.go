package analytics

import "time"

func mapNotNull[T any, U any](v *T, f func(T) U) *U {
	if v == nil {
		return nil
	}

	mapped := f(*v)
	return &mapped
}

func mapNullableSecondsToDuration(seconds *int64) *time.Duration {
	return mapNotNull(seconds, func(secs int64) time.Duration {
		return time.Duration(secs) * time.Second
	})
}
