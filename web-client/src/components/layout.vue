<script setup lang="ts">
import myHeader from "./header/my-header.vue";
</script>

<template>
	<myHeader />
	<div class="content-wrapper">
		<slot />
	</div>
</template>

<style>
.content-wrapper {
	max-width: 1024px;
	margin: 0 auto;
}
</style>
