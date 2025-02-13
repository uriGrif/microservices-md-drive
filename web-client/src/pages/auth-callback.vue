<script setup lang="ts">
import { useAuth0 } from "@auth0/auth0-vue";
import { useRouter } from "vue-router";

const { error } = useAuth0();

if (!error) {
	const router = useRouter();
	router.push("/");
}
</script>

<template>
	<div v-if="error">
		<h1>Error</h1>
		<h2>{{ error.message }}</h2>
	</div>
</template>
