package helpers

func Pipe2[T0, T1, T2 any](v T0, f1 F1[T0, T1], f2 F1[T1, T2]) T2 {
	return f2(f1(v))
}

func Pipe3[T0, T1, T2, T3 any](v T0, f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3]) T3 {
	return f3(f2(f1(v)))
}

func Pipe4[T0, T1, T2, T3, T4 any](v T0, f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4]) T4 {
	return f4(f3(f2(f1(v))))
}

func Pipe5[T0, T1, T2, T3, T4, T5 any](v T0, f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5]) T5 {
	return f5(f4(f3(f2(f1(v)))))
}

func Pipe6[T0, T1, T2, T3, T4, T5, T6 any](v T0, f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6]) T6 {
	return f6(f5(f4(f3(f2(f1(v))))))
}

func Pipe7[T0, T1, T2, T3, T4, T5, T6, T7 any](v T0, f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7]) T7 {
	return f7(f6(f5(f4(f3(f2(f1(v)))))))
}

func Pipe8[T0, T1, T2, T3, T4, T5, T6, T7, T8 any](v T0, f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7], f8 F1[T7, T8]) T8 {
	return f8(f7(f6(f5(f4(f3(f2(f1(v))))))))
}

func Pipe9[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v T0, f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7], f8 F1[T7, T8], f9 F1[T8, T9]) T9 {
	return f9(f8(f7(f6(f5(f4(f3(f2(f1(v)))))))))
}

func Pipe10[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](v T0, f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7], f8 F1[T7, T8], f9 F1[T8, T9], f10 F1[T9, T10]) T10 {
	return f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(v))))))))))
}

func Pipe11[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any](v T0, f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7], f8 F1[T7, T8], f9 F1[T8, T9], f10 F1[T9, T10], f11 F1[T10, T11]) T11 {
	return f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(v)))))))))))
}

func Pipe12[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any](v T0, f1 F1[T0, T1], f2 F1[T1, T2], f3 F1[T2, T3], f4 F1[T3, T4], f5 F1[T4, T5], f6 F1[T5, T6], f7 F1[T6, T7], f8 F1[T7, T8], f9 F1[T8, T9], f10 F1[T9, T10], f11 F1[T10, T11], f12 F1[T11, T12]) T12 {
	return f12(f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(v))))))))))))
}
