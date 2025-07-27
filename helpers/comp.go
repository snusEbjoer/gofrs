package helpers

func Comp2[T0, T1, T2 any](f1 F1[T0, T1], f2 F1[T1, T2]) func(T0) T2 {
	return func(v T0) T2 {
		return f2(f1(v))
	}
}

func Comp3[T0, T1, T2, T3 any](f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3]) func(T0) T3 {
	return func(v T0) T3 {
		return f3(f2(f1(v)))
	}
}

func Comp4[T0, T1, T2, T3, T4 any](f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4]) func(T0) T4 {
	return func(v T0) T4 {
		return f4(f3(f2(f1(v))))
	}
}

func Comp5[T0, T1, T2, T3, T4, T5 any](f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5]) func(T0) T5 {
	return func(v T0) T5 {
		return f5(f4(f3(f2(f1(v)))))
	}
}

func Comp6[T0, T1, T2, T3, T4, T5, T6 any](f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6]) func(T0) T6 {
	return func(v T0) T6 {
		return f6(f5(f4(f3(f2(f1(v))))))
	}
}

func Comp7[T0, T1, T2, T3, T4, T5, T6, T7 any](f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7]) func(T0) T7 {
	return func(v T0) T7 {
		return f7(f6(f5(f4(f3(f2(f1(v)))))))
	}
}

func Comp8[T0, T1, T2, T3, T4, T5, T6, T7, T8 any](f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7], f8 F1[T7, T8]) func(T0) T8 {
	return func(v T0) T8 {
		return f8(f7(f6(f5(f4(f3(f2(f1(v))))))))
	}
}

func Comp9[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9 any](f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7], f8 F1[T7, T8], f9 F1[T8, T9]) func(T0) T9 {
	return func(v T0) T9 {
		return f9(f8(f7(f6(f5(f4(f3(f2(f1(v)))))))))
	}
}

func Comp10[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7], f8 F1[T7, T8], f9 F1[T8, T9], f10 F1[T9, T10]) func(T0) T10 {
	return func(v T0) T10 {
		return f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(v))))))))))
	}
}

func Comp11[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any](f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7], f8 F1[T7, T8], f9 F1[T8, T9], f10 F1[T9, T10], f11 F1[T10, T11]) func(T0) T11 {
	return func(v T0) T11 {
		return f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(v)))))))))))
	}
}

func Comp12[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any](f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7], f8 F1[T7, T8], f9 F1[T8, T9], f10 F1[T9, T10], f11 F1[T10, T11], f12 F1[T11, T12]) func(T0) T12 {
	return func(v T0) T12 {
		return f12(f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(v))))))))))))
	}
}
