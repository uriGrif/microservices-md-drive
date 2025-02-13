<script setup lang="ts">
import { onMounted, ref } from "vue";
import layout from "../components/layout.vue";
import { useAuth0 } from "@auth0/auth0-vue";

const { user, getAccessTokenSilently } = useAuth0();

const token = ref<string | null>(null);

onMounted(async () => {
	try {
		token.value = await getAccessTokenSilently();
	} catch (error) {
		console.error("Error fetching token:", error);
	}
});
</script>

<template>
	<layout>
		Esta es la home

		<div v-if="user">
			<p>
				{{ token }}
			</p>
			<p>
				{{ JSON.stringify(user, null, 1) }}
			</p>
		</div>
	</layout>
</template>
