<script lang="ts">
    import SunIcon from "@lucide/svelte/icons/sun";
    import MoonIcon from "@lucide/svelte/icons/moon";
    
    import { resetMode, setMode } from "mode-watcher";
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
    import { buttonVariants } from "$lib/components/ui/button/index.js";
    import { userPrefersMode } from "mode-watcher";

</script>
 
<DropdownMenu.Root>
 <DropdownMenu.Trigger
  class={buttonVariants({ variant: "outline", size: "icon" })}
 >
  <SunIcon
   class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 !transition-all dark:-rotate-90 dark:scale-0"
  />
  <MoonIcon
   class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 !transition-all dark:rotate-0 dark:scale-100"
  />
  <span class="sr-only">Toggle theme</span>
 </DropdownMenu.Trigger>
 <DropdownMenu.Content align="end">
  <DropdownMenu.Item class={userPrefersMode.current === "light" ? "active-mode-item": ""} onclick={() => setMode("light")}>Light</DropdownMenu.Item>
  <DropdownMenu.Item class={userPrefersMode.current === "dark" ? "active-mode-item": ""} onclick={() => setMode("dark")}>Dark</DropdownMenu.Item>
  <DropdownMenu.Item class={userPrefersMode.current === "system" ? "active-mode-item": ""} onclick={() => resetMode()}>System</DropdownMenu.Item>
 </DropdownMenu.Content>
</DropdownMenu.Root>

<style>
    :global(.active-mode-item) {
        font-weight: bold;
        background-color: var(--muted);
    }
</style>