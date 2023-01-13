package bootstrap

func Init() {

	Cfg._init()
	Log._init()
	DB._init()
	AsyNq._init()
	Router._init()
}
