/*
*	 ____  _   _  _____        _______ _        _    _  _______
*	/ ___|| \ | |/ _ \ \      / /  ___| |      / \  | |/ / ____|
*	\___ \|  \| | | | \ \ /\ / /| |_  | |     / _ \ | ' /|  _|
*	 ___) | |\  | |_| |\ V  V / |  _| | |___ / ___ \| . \| |___
*	|____/|_| \_|\___/  \_/\_/  |_|   |_____/_/   \_\_|\_\_____|
*
*  https://github.com/yingzhuo/snowflake
 */
package prome

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	IdCreatedCounter *prometheus.CounterVec
)

func init() {
	IdCreatedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "id_created",
			Help: "Number of uuid created.",
		},
		[]string{"app"},
	)

	prometheus.MustRegister(IdCreatedCounter)
}
