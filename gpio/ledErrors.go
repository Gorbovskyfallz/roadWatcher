package gpio

import (
	"kek/flash"
	"kek/network"
)

//индикация отсутствия интернета
// индикация отсутствия впн
// индикация ошибки потока
// индикация ошибки монтирования
// индикация "все хорошо"

// сначала мы накидаем просто названивая функций, надо попробовать развивать мы
//шление относительно уровней абстракции

type LedIndicator interface {
	NetworkIndication(Network network.Network)
	SystemIndication(StatFlash flash.StatFlash)
}

func (i *IoPins) NetworkIndication(Network network.Network) {

}

func (i *IoPins) SystemIndication(StatFlash flash.StatFlash) {

}
