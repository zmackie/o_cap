// Code generated by "stringer -type=CAP"; DO NOT EDIT.

package main

import "strconv"

const _CAP_name = "CAP_CHOWNCAP_DAC_OVERRIDECAP_DAC_READ_SEARCHCAP_FOWNERCAP_FSETIDCAP_KILLCAP_SETGIDCAP_SETUIDCAP_SETPCAPCAP_LINX_IMMUTABLECAP_NET_BIND_SERVICECAP_NET_BROADCASTCAP_NET_ADMINCAP_NET_RAWCAP_IPC_LOCKCAP_IPC_OWNERCAP_SYS_MODULECAP_SYS_RAWIOCAP_SYS_CHROOTCAP_SYS_PTRACECAP_SYS_PACCTCAP_SYS_ADMINCAP_SYS_BOOTCAP_SYS_NICECAP_SYS_RESOURCECAP_SYS_TIMECAP_SYS_TTY_CONFIGCAP_MKNODCAP_LEASECAP_AUDIT_WRITECAP_AUDIT_CONTROLCAP_SETFCAPCAP_MAC_OVERRIDECAP_MAC_ADMINCAP_SYSLOGCAP_WAKE_ALARMCAP_BLOCK_SUSPENDCAP_AUDIT_READ"

var _CAP_index = [...]uint16{0, 9, 25, 44, 54, 64, 72, 82, 92, 103, 121, 141, 158, 171, 182, 194, 207, 221, 234, 248, 262, 275, 288, 300, 312, 328, 340, 358, 367, 376, 391, 408, 419, 435, 448, 458, 472, 489, 503}

func (i CAP) String() string {
	if i < 0 || i >= CAP(len(_CAP_index)-1) {
		return "CAP(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CAP_name[_CAP_index[i]:_CAP_index[i+1]]
}
