<script lang="ts">
	import Button from '$components/buttons/Button.svelte';
	import IconWarning from '$components/icons/IconWarning.svelte';
	import { env } from '$env/dynamic/public';
	import LL from '$i18n/i18n-svelte';
	import { expandCollapse } from '$ts/animation/transitions';
	import {
		currentServer,
		currentServerHealthStatus,
		defaultServer,
		defaultServerHealthStatus
	} from '$ts/stores/serverHealth';
	import { serverUrl } from '$ts/stores/serverUrl';

	const switchToDefaultServer = () => {
		serverUrl.set(env.PUBLIC_DEFAULT_SERVER_URL);
		localStorage.removeItem('serverUrl');
		currentServer.set($defaultServer);
		currentServerHealthStatus.set($defaultServerHealthStatus);
	};
</script>

<div transition:expandCollapse|local={{ duration: 300 }} class="relative">
	<div class="pb-7 flex flex-col items-center">
		<div
			class="py-4 px-4 text-xs md:text-sm gap-3 flex items-center justify-start max-w-[34rem] rounded-xl bg-c-danger/8 text-c-danger"
		>
			<IconWarning class="w-6 h-6" />
			<p class="flex-1">
				{$LL.Error.ServerSeemsOffline()}
			</p>
		</div>
		{#if env.PUBLIC_DEFAULT_SERVER_URL && $serverUrl !== env.PUBLIC_DEFAULT_SERVER_URL && ($defaultServerHealthStatus === 'healthy' || $defaultServerHealthStatus === 'loading')}
			<div transition:expandCollapse|local={{ duration: 300 }} class="relative">
				<div class="pt-3.5">
					<Button size="sm" onClick={switchToDefaultServer}>
						{$LL.Shared.SwitchToDefaultServerButton()}
					</Button>
				</div>
			</div>
		{/if}
	</div>
</div>
