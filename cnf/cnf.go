/*
*	 ____  _   _  _____        _______ _        _    _  _______
*	/ ___|| \ | |/ _ \ \      / /  ___| |      / \  | |/ / ____|
*	\___ \|  \| | | | \ \ /\ / /| |_  | |     / _ \ | ' /|  _|
*	 ___) | |\  | |_| |\ V  V / |  _| | |___ / ___ \| . \| |___
*	|____/|_| \_|\___/  \_/\_/  |_|   |_____/_/   \_\_|\_\_____|
*
*  https://github.com/yingzhuo/snowflake
 */
package cnf

import "github.com/bwmarrin/snowflake"

var (
	NodeId        int64
	Port          int
	Indent        bool
	SnowflakeNode *snowflake.Node
)
