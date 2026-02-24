<script setup lang="ts">
import { RouterView } from "vue-router";
import { NMessageProvider, NModalProvider, NConfigProvider } from "naive-ui";
import themeOverrides from "./shared/theme-overrides";
</script>

<template>
	<n-config-provider :theme-overrides="themeOverrides">
		<n-message-provider>
			<n-modal-provider>
				<RouterView />
			</n-modal-provider>
		</n-message-provider>
	</n-config-provider>
</template>

<style></style>
