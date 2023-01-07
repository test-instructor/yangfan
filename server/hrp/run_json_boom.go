package hrp

func (b *HRPBoomer) InitBoomerCheetah() {
	b.SetSpawnCount(b.GetProfile().SpawnCount)
	b.SetSpawnRate(b.GetProfile().SpawnRate)
	b.SetRunTime(b.GetProfile().RunTime)
	if b.GetProfile().LoopCount > 0 {
		b.SetLoopCount(b.GetProfile().LoopCount)
	}
	b.SetRateLimiter(b.GetProfile().MaxRPS, b.GetProfile().RequestIncreaseRate)
	b.SetDisableKeepAlive(b.GetProfile().DisableKeepalive)
	b.SetDisableCompression(b.GetProfile().DisableCompression)
	b.SetClientTransport()
	b.EnableCPUProfile(b.GetProfile().CPUProfile, b.GetProfile().CPUProfileDuration)
	b.EnableMemoryProfile(b.GetProfile().MemoryProfile, b.GetProfile().MemoryProfileDuration)
}
