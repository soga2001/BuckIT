<script lang="ts">
	import ThemeSelector from "$lib/components/ui/theme-selector.svelte";
	import { userPrefersMode } from "mode-watcher";
	import "../app.css";
	import { ModeWatcher } from "mode-watcher";
	let { children } = $props();
</script>

<ThemeSelector />
<ModeWatcher defaultMode="system" />
{@render children?.()}