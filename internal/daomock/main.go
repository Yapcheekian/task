package daomock

//go:generate mockgen -destination=mock.go -package=$GOPACKAGE github.com/Yapcheekian/task/pkg/dao TaskDAO
