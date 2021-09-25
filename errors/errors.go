package errors

type GMOPSError interface {
	Error() string
	Message() string
	Code() string
	CanRetry() bool
}

// NewErrors converts GMO-PS errors to Struct
// Examples:
//		["CT0009", "CT0035"]
func NewErrors(errInfos ...string) []GMOPSError {
	errs := make([]GMOPSError, 0, len(errInfos))
	for _, info := range errInfos {
		switch info {
		case "SP0001":
			errs = append(errs, &PS_SP0001{})
		case "SP0002":
			errs = append(errs, &PS_SP0002{})
		case "SP0003":
			errs = append(errs, &PS_SP0003{})
		case "SP0004":
			errs = append(errs, &PS_SP0004{})
		case "SP0005":
			errs = append(errs, &PS_SP0005{})
		case "SP0006":
			errs = append(errs, &PS_SP0006{})
		case "SP0007":
			errs = append(errs, &PS_SP0007{})
		case "SP0008":
			errs = append(errs, &PS_SP0008{})
		case "SP0009":
			errs = append(errs, &PS_SP0009{})
		case "SP0010":
			errs = append(errs, &PS_SP0010{})
		case "SP0011":
			errs = append(errs, &PS_SP0011{})
		case "SP0012":
			errs = append(errs, &PS_SP0012{})
		case "SP0013":
			errs = append(errs, &PS_SP0013{})
		case "SP0014":
			errs = append(errs, &PS_SP0014{})
		case "SP0015":
			errs = append(errs, &PS_SP0015{})
		case "SP0016":
			errs = append(errs, &PS_SP0016{})
		case "SP0017":
			errs = append(errs, &PS_SP0017{})
		case "SP0018":
			errs = append(errs, &PS_SP0018{})
		case "SP0019":
			errs = append(errs, &PS_SP0019{})
		case "CT0001":
			errs = append(errs, &PS_CT0001{})
		case "CT0002":
			errs = append(errs, &PS_CT0002{})
		case "CT0003":
			errs = append(errs, &PS_CT0003{})
		case "CT0004":
			errs = append(errs, &PS_CT0004{})
		case "CT0005":
			errs = append(errs, &PS_CT0005{})
		case "CT0006":
			errs = append(errs, &PS_CT0006{})
		case "CT0007":
			errs = append(errs, &PS_CT0007{})
		case "CT0008":
			errs = append(errs, &PS_CT0008{})
		case "CT0009":
			errs = append(errs, &PS_CT0009{})
		case "CT0010":
			errs = append(errs, &PS_CT0010{})
		case "CT0011":
			errs = append(errs, &PS_CT0011{})
		case "CT0012":
			errs = append(errs, &PS_CT0012{})
		case "CT0013":
			errs = append(errs, &PS_CT0013{})
		case "CT0014":
			errs = append(errs, &PS_CT0014{})
		case "CT0015":
			errs = append(errs, &PS_CT0015{})
		case "CT0016":
			errs = append(errs, &PS_CT0016{})
		case "CT0017":
			errs = append(errs, &PS_CT0017{})
		case "CT0018":
			errs = append(errs, &PS_CT0018{})
		case "CT0019":
			errs = append(errs, &PS_CT0019{})
		case "CT0020":
			errs = append(errs, &PS_CT0020{})
		case "CT0021":
			errs = append(errs, &PS_CT0021{})
		case "CT0022":
			errs = append(errs, &PS_CT0022{})
		case "CT0023":
			errs = append(errs, &PS_CT0023{})
		case "CT0024":
			errs = append(errs, &PS_CT0024{})
		case "CT0025":
			errs = append(errs, &PS_CT0025{})
		case "CT0026":
			errs = append(errs, &PS_CT0026{})
		case "CT0027":
			errs = append(errs, &PS_CT0027{})
		case "CT0028":
			errs = append(errs, &PS_CT0028{})
		case "CT0029":
			errs = append(errs, &PS_CT0029{})
		case "CT0030":
			errs = append(errs, &PS_CT0030{})
		case "CT0031":
			errs = append(errs, &PS_CT0031{})
		case "CT0032":
			errs = append(errs, &PS_CT0032{})
		case "CT0033":
			errs = append(errs, &PS_CT0033{})
		case "CT0034":
			errs = append(errs, &PS_CT0034{})
		case "CT0035":
			errs = append(errs, &PS_CT0035{})
		case "CT0036":
			errs = append(errs, &PS_CT0036{})
		case "CT0037":
			errs = append(errs, &PS_CT0037{})
		case "CT0038":
			errs = append(errs, &PS_CT0038{})
		case "CT0039":
			errs = append(errs, &PS_CT0039{})
		case "CT0040":
			errs = append(errs, &PS_CT0040{})
		case "CT0041":
			errs = append(errs, &PS_CT0041{})
		case "CT0042":
			errs = append(errs, &PS_CT0042{})
		case "CT0043":
			errs = append(errs, &PS_CT0043{})
		case "CT0044":
			errs = append(errs, &PS_CT0044{})
		case "CT0045":
			errs = append(errs, &PS_CT0045{})
		case "CT0046":
			errs = append(errs, &PS_CT0046{})
		case "CT0047":
			errs = append(errs, &PS_CT0047{})
		case "CT0048":
			errs = append(errs, &PS_CT0048{})
		case "CT0049":
			errs = append(errs, &PS_CT0049{})
		case "CT0050":
			errs = append(errs, &PS_CT0050{})
		case "CT0051":
			errs = append(errs, &PS_CT0051{})
		case "CT0052":
			errs = append(errs, &PS_CT0052{})
		case "CT0053":
			errs = append(errs, &PS_CT0053{})
		case "CT0054":
			errs = append(errs, &PS_CT0054{})
		case "CT0055":
			errs = append(errs, &PS_CT0055{})
		case "CT0056":
			errs = append(errs, &PS_CT0056{})
		case "CT0057":
			errs = append(errs, &PS_CT0057{})
		case "CT0059":
			errs = append(errs, &PS_CT0059{})
		case "CT0060":
			errs = append(errs, &PS_CT0060{})
		case "CT0061":
			errs = append(errs, &PS_CT0061{})
		case "CT0062":
			errs = append(errs, &PS_CT0062{})
		case "CT0063":
			errs = append(errs, &PS_CT0063{})
		case "CT0064":
			errs = append(errs, &PS_CT0064{})
		case "CT0065":
			errs = append(errs, &PS_CT0065{})
		case "CT0066":
			errs = append(errs, &PS_CT0066{})
		case "CT0075":
			errs = append(errs, &PS_CT0075{})
		case "CT0076":
			errs = append(errs, &PS_CT0076{})
		case "CT0077":
			errs = append(errs, &PS_CT0077{})
		case "CT0078":
			errs = append(errs, &PS_CT0078{})
		case "DL0001":
			errs = append(errs, &PS_DL0001{})
		case "DL0002":
			errs = append(errs, &PS_DL0002{})
		case "DL0003":
			errs = append(errs, &PS_DL0003{})
		case "DL0004":
			errs = append(errs, &PS_DL0004{})
		case "DL0005":
			errs = append(errs, &PS_DL0005{})
		case "DL0006":
			errs = append(errs, &PS_DL0006{})
		case "DL0007":
			errs = append(errs, &PS_DL0007{})
		case "DL0008":
			errs = append(errs, &PS_DL0008{})
		case "DL0009":
			errs = append(errs, &PS_DL0009{})
		case "DL0010":
			errs = append(errs, &PS_DL0010{})
		case "DL0011":
			errs = append(errs, &PS_DL0011{})
		case "DL0012":
			errs = append(errs, &PS_DL0012{})
		case "DL0013":
			errs = append(errs, &PS_DL0013{})
		case "DL0014":
			errs = append(errs, &PS_DL0014{})
		case "DL0015":
			errs = append(errs, &PS_DL0015{})
		case "DL0016":
			errs = append(errs, &PS_DL0016{})
		case "DL0017":
			errs = append(errs, &PS_DL0017{})
		case "DL0018":
			errs = append(errs, &PS_DL0018{})
		case "DL0019":
			errs = append(errs, &PS_DL0019{})
		case "DL0020":
			errs = append(errs, &PS_DL0020{})
		case "DL0021":
			errs = append(errs, &PS_DL0021{})
		case "DT0001":
			errs = append(errs, &PS_DT0001{})
		case "DT0002":
			errs = append(errs, &PS_DT0002{})
		case "DT0003":
			errs = append(errs, &PS_DT0003{})
		case "DT0004":
			errs = append(errs, &PS_DT0004{})
		case "DT0005":
			errs = append(errs, &PS_DT0005{})
		case "DT0006":
			errs = append(errs, &PS_DT0006{})
		case "DT0007":
			errs = append(errs, &PS_DT0007{})
		case "DT0008":
			errs = append(errs, &PS_DT0008{})
		case "DT0009":
			errs = append(errs, &PS_DT0009{})
		case "DT0010":
			errs = append(errs, &PS_DT0010{})
		case "DT0014":
			errs = append(errs, &PS_DT0014{})
		case "DT0015":
			errs = append(errs, &PS_DT0015{})
		case "DT0016":
			errs = append(errs, &PS_DT0016{})
		case "DT0017":
			errs = append(errs, &PS_DT0017{})
		case "GT0001":
			errs = append(errs, &PS_GT0001{})
		case "GT0002":
			errs = append(errs, &PS_GT0002{})
		case "GT0003":
			errs = append(errs, &PS_GT0003{})
		case "GT0004":
			errs = append(errs, &PS_GT0004{})
		case "GT0005":
			errs = append(errs, &PS_GT0005{})
		case "GT0006":
			errs = append(errs, &PS_GT0006{})
		case "GT0007":
			errs = append(errs, &PS_GT0007{})
		case "GT0008":
			errs = append(errs, &PS_GT0008{})
		case "GT0009":
			errs = append(errs, &PS_GT0009{})
		case "ST0001":
			errs = append(errs, &PS_ST0001{})
		case "ST0002":
			errs = append(errs, &PS_ST0002{})

		}
	}
	return errs
}
