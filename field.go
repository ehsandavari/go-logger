package logger

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

type IField interface {
	WithBinary(key string, value []byte) ILogger
	WithBool(key string, value bool) ILogger
	WithBoolp(key string, value *bool) ILogger
	WithByteString(key string, value []byte) ILogger
	WithComplex128(key string, value complex128) ILogger
	WithComplex128p(key string, value *complex128) ILogger
	WithComplex64(key string, value complex64) ILogger
	WithComplex64p(key string, value *complex64) ILogger
	WithFloat64(key string, value float64) ILogger
	WithFloat64p(key string, value *float64) ILogger
	WithFloat32(key string, value float32) ILogger
	WithFloat32p(key string, value *float32) ILogger
	WithInt(key string, value int) ILogger
	WithIntp(key string, value *int) ILogger
	WithInt64(key string, value int64) ILogger
	WithInt64p(key string, value *int64) ILogger
	WithInt32(key string, value int32) ILogger
	WithInt32p(key string, value *int32) ILogger
	WithInt16(key string, value int16) ILogger
	WithInt16p(key string, value *int16) ILogger
	WithInt8(key string, value int8) ILogger
	WithInt8p(key string, value *int8) ILogger
	WithUint(key string, value uint) ILogger
	WithUintp(key string, value *uint) ILogger
	WithUint64(key string, value uint64) ILogger
	WithUint64p(key string, value *uint64) ILogger
	WithUint32(key string, value uint32) ILogger
	WithUint32p(key string, value *uint32) ILogger
	WithUint16(key string, value uint16) ILogger
	WithUint16p(key string, value *uint16) ILogger
	WithUint8(key string, value uint8) ILogger
	WithUint8p(key string, value *uint8) ILogger
	WithUintptr(key string, value uintptr) ILogger
	WithUintptrp(key string, value *uintptr) ILogger
	WithReflect(key string, value interface{}) ILogger
	WithNamespace(key string) ILogger
	WithStringer(key string, value fmt.Stringer) ILogger
	WithString(key string, value string) ILogger
	WithStringp(key string, value *string) ILogger
	WithTime(key string, value time.Time) ILogger
	WithTimep(key string, value *time.Time) ILogger
	WithStack(key string) ILogger
	WithStackSkip(key string, skip int) ILogger
	WithDuration(key string, value time.Duration) ILogger
	WithDurationp(key string, value *time.Duration) ILogger
	WithAny(key string, value interface{}) ILogger
	WithError(value error) ILogger
	WithNamedError(key string, value error) ILogger
	WithBools(key string, value []bool) ILogger
	WithByteStrings(key string, value [][]byte) ILogger
	WithComplex128s(key string, value []complex128) ILogger
	WithComplex64s(key string, value []complex64) ILogger
	WithDurations(key string, value []time.Duration) ILogger
	WithFloat64s(key string, value []float64) ILogger
	WithFloat32s(key string, value []float32) ILogger
	WithInts(key string, value []int) ILogger
	WithInt64s(key string, value []int64) ILogger
	WithInt32s(key string, value []int32) ILogger
	WithInt16s(key string, value []int16) ILogger
	WithInt8s(key string, value []int8) ILogger
	WithStrings(key string, value []string) ILogger
	WithTimes(key string, value []time.Time) ILogger
	WithUints(key string, value []uint) ILogger
	WithUint64s(key string, value []uint64) ILogger
	WithUint32s(key string, value []uint32) ILogger
	WithUint16s(key string, value []uint16) ILogger
	WithUint8s(key string, value []uint8) ILogger
	WithUintptrs(key string, value []uintptr) ILogger
	WithErrors(key string, value []error) ILogger
	WithStringers(key string, value []fmt.Stringer) ILogger
}

func (r *sLogger) WithBinary(key string, value []byte) ILogger {
	r.fields = append(r.fields, zap.Binary(key, value))
	return r
}

func (r *sLogger) WithBool(key string, value bool) ILogger {
	r.fields = append(r.fields, zap.Bool(key, value))
	return r
}

func (r *sLogger) WithBoolp(key string, value *bool) ILogger {
	r.fields = append(r.fields, zap.Boolp(key, value))
	return r
}

func (r *sLogger) WithByteString(key string, value []byte) ILogger {
	r.fields = append(r.fields, zap.ByteString(key, value))
	return r
}

func (r *sLogger) WithComplex128(key string, value complex128) ILogger {
	r.fields = append(r.fields, zap.Complex128(key, value))
	return r
}

func (r *sLogger) WithComplex128p(key string, value *complex128) ILogger {
	r.fields = append(r.fields, zap.Complex128p(key, value))
	return r
}

func (r *sLogger) WithComplex64(key string, value complex64) ILogger {
	r.fields = append(r.fields, zap.Complex64(key, value))
	return r
}

func (r *sLogger) WithComplex64p(key string, value *complex64) ILogger {
	r.fields = append(r.fields, zap.Complex64p(key, value))
	return r
}

func (r *sLogger) WithFloat64(key string, value float64) ILogger {
	r.fields = append(r.fields, zap.Float64(key, value))
	return r
}

func (r *sLogger) WithFloat64p(key string, value *float64) ILogger {
	r.fields = append(r.fields, zap.Float64p(key, value))
	return r
}

func (r *sLogger) WithFloat32(key string, value float32) ILogger {
	r.fields = append(r.fields, zap.Float32(key, value))
	return r
}

func (r *sLogger) WithFloat32p(key string, value *float32) ILogger {
	r.fields = append(r.fields, zap.Float32p(key, value))
	return r
}

func (r *sLogger) WithInt(key string, value int) ILogger {
	r.fields = append(r.fields, zap.Int(key, value))
	return r
}

func (r *sLogger) WithIntp(key string, value *int) ILogger {
	r.fields = append(r.fields, zap.Intp(key, value))
	return r
}

func (r *sLogger) WithInt64(key string, value int64) ILogger {
	r.fields = append(r.fields, zap.Int64(key, value))
	return r
}

func (r *sLogger) WithInt64p(key string, value *int64) ILogger {
	r.fields = append(r.fields, zap.Int64p(key, value))
	return r
}

func (r *sLogger) WithInt32(key string, value int32) ILogger {
	r.fields = append(r.fields, zap.Int32(key, value))
	return r
}

func (r *sLogger) WithInt32p(key string, value *int32) ILogger {
	r.fields = append(r.fields, zap.Int32p(key, value))
	return r
}

func (r *sLogger) WithInt16(key string, value int16) ILogger {
	r.fields = append(r.fields, zap.Int16(key, value))
	return r
}

func (r *sLogger) WithInt16p(key string, value *int16) ILogger {
	r.fields = append(r.fields, zap.Int16p(key, value))
	return r
}

func (r *sLogger) WithInt8(key string, value int8) ILogger {
	r.fields = append(r.fields, zap.Int8(key, value))
	return r
}

func (r *sLogger) WithInt8p(key string, value *int8) ILogger {
	r.fields = append(r.fields, zap.Int8p(key, value))
	return r
}

func (r *sLogger) WithUint(key string, value uint) ILogger {
	r.fields = append(r.fields, zap.Uint(key, value))
	return r
}

func (r *sLogger) WithUintp(key string, value *uint) ILogger {
	r.fields = append(r.fields, zap.Uintp(key, value))
	return r
}

func (r *sLogger) WithUint64(key string, value uint64) ILogger {
	r.fields = append(r.fields, zap.Uint64(key, value))
	return r
}

func (r *sLogger) WithUint64p(key string, value *uint64) ILogger {
	r.fields = append(r.fields, zap.Uint64p(key, value))
	return r
}

func (r *sLogger) WithUint32(key string, value uint32) ILogger {
	r.fields = append(r.fields, zap.Uint32(key, value))
	return r
}

func (r *sLogger) WithUint32p(key string, value *uint32) ILogger {
	r.fields = append(r.fields, zap.Uint32p(key, value))
	return r
}

func (r *sLogger) WithUint16(key string, value uint16) ILogger {
	r.fields = append(r.fields, zap.Uint16(key, value))
	return r
}

func (r *sLogger) WithUint16p(key string, value *uint16) ILogger {
	r.fields = append(r.fields, zap.Uint16p(key, value))
	return r
}

func (r *sLogger) WithUint8(key string, value uint8) ILogger {
	r.fields = append(r.fields, zap.Uint8(key, value))
	return r
}

func (r *sLogger) WithUint8p(key string, value *uint8) ILogger {
	r.fields = append(r.fields, zap.Uint8p(key, value))
	return r
}

func (r *sLogger) WithUintptr(key string, value uintptr) ILogger {
	r.fields = append(r.fields, zap.Uintptr(key, value))
	return r
}

func (r *sLogger) WithUintptrp(key string, value *uintptr) ILogger {
	r.fields = append(r.fields, zap.Uintptrp(key, value))
	return r
}

func (r *sLogger) WithReflect(key string, value any) ILogger {
	r.fields = append(r.fields, zap.Reflect(key, value))
	return r
}

func (r *sLogger) WithNamespace(key string) ILogger {
	r.fields = append(r.fields, zap.Namespace(key))
	return r
}

func (r *sLogger) WithStringer(key string, value fmt.Stringer) ILogger {
	r.fields = append(r.fields, zap.Stringer(key, value))
	return r
}

func (r *sLogger) WithString(key string, value string) ILogger {
	r.fields = append(r.fields, zap.String(key, value))
	return r
}

func (r *sLogger) WithStringp(key string, value *string) ILogger {
	r.fields = append(r.fields, zap.Stringp(key, value))
	return r
}

func (r *sLogger) WithTime(key string, value time.Time) ILogger {
	r.fields = append(r.fields, zap.Time(key, value))
	return r
}

func (r *sLogger) WithTimep(key string, value *time.Time) ILogger {
	r.fields = append(r.fields, zap.Timep(key, value))
	return r
}

func (r *sLogger) WithStack(key string) ILogger {
	r.fields = append(r.fields, zap.Stack(key))
	return r
}

func (r *sLogger) WithStackSkip(key string, skip int) ILogger {
	r.fields = append(r.fields, zap.StackSkip(key, skip))
	return r
}

func (r *sLogger) WithDuration(key string, value time.Duration) ILogger {
	r.fields = append(r.fields, zap.Duration(key, value))
	return r
}

func (r *sLogger) WithDurationp(key string, value *time.Duration) ILogger {
	r.fields = append(r.fields, zap.Durationp(key, value))
	return r
}

func (r *sLogger) WithAny(key string, value any) ILogger {
	r.fields = append(r.fields, zap.Any(key, value))
	return r
}

func (r *sLogger) WithError(value error) ILogger {
	r.fields = append(r.fields, zap.Error(value))
	return r
}

func (r *sLogger) WithNamedError(key string, value error) ILogger {
	r.fields = append(r.fields, zap.NamedError(key, value))
	return r
}

func (r *sLogger) WithBools(key string, value []bool) ILogger {
	r.fields = append(r.fields, zap.Bools(key, value))
	return r
}

func (r *sLogger) WithByteStrings(key string, value [][]byte) ILogger {
	r.fields = append(r.fields, zap.ByteStrings(key, value))
	return r
}

func (r *sLogger) WithComplex128s(key string, value []complex128) ILogger {
	r.fields = append(r.fields, zap.Complex128s(key, value))
	return r
}

func (r *sLogger) WithComplex64s(key string, value []complex64) ILogger {
	r.fields = append(r.fields, zap.Complex64s(key, value))
	return r
}

func (r *sLogger) WithDurations(key string, value []time.Duration) ILogger {
	r.fields = append(r.fields, zap.Durations(key, value))
	return r
}

func (r *sLogger) WithFloat64s(key string, value []float64) ILogger {
	r.fields = append(r.fields, zap.Float64s(key, value))
	return r
}

func (r *sLogger) WithFloat32s(key string, value []float32) ILogger {
	r.fields = append(r.fields, zap.Float32s(key, value))
	return r
}

func (r *sLogger) WithInts(key string, value []int) ILogger {
	r.fields = append(r.fields, zap.Ints(key, value))
	return r
}

func (r *sLogger) WithInt64s(key string, value []int64) ILogger {
	r.fields = append(r.fields, zap.Int64s(key, value))
	return r
}

func (r *sLogger) WithInt32s(key string, value []int32) ILogger {
	r.fields = append(r.fields, zap.Int32s(key, value))
	return r
}

func (r *sLogger) WithInt16s(key string, value []int16) ILogger {
	r.fields = append(r.fields, zap.Int16s(key, value))
	return r
}

func (r *sLogger) WithInt8s(key string, value []int8) ILogger {
	r.fields = append(r.fields, zap.Int8s(key, value))
	return r
}

func (r *sLogger) WithStrings(key string, value []string) ILogger {
	r.fields = append(r.fields, zap.Strings(key, value))
	return r
}

func (r *sLogger) WithTimes(key string, value []time.Time) ILogger {
	r.fields = append(r.fields, zap.Times(key, value))
	return r
}

func (r *sLogger) WithUints(key string, value []uint) ILogger {
	r.fields = append(r.fields, zap.Uints(key, value))
	return r
}

func (r *sLogger) WithUint64s(key string, value []uint64) ILogger {
	r.fields = append(r.fields, zap.Uint64s(key, value))
	return r
}

func (r *sLogger) WithUint32s(key string, value []uint32) ILogger {
	r.fields = append(r.fields, zap.Uint32s(key, value))
	return r
}

func (r *sLogger) WithUint16s(key string, value []uint16) ILogger {
	r.fields = append(r.fields, zap.Uint16s(key, value))
	return r
}

func (r *sLogger) WithUint8s(key string, value []uint8) ILogger {
	r.fields = append(r.fields, zap.Uint8s(key, value))
	return r
}

func (r *sLogger) WithUintptrs(key string, value []uintptr) ILogger {
	r.fields = append(r.fields, zap.Uintptrs(key, value))
	return r
}

func (r *sLogger) WithErrors(key string, value []error) ILogger {
	r.fields = append(r.fields, zap.Errors(key, value))
	return r
}

func (r *sLogger) WithStringers(key string, value []fmt.Stringer) ILogger {
	r.fields = append(r.fields, zap.Stringers(key, value))
	return r
}
