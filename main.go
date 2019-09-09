package main

import "calculabotv2/common"

func main() {
	NewBot(common.GetConfig().GetToken()).Start()
}
