<script setup lang="ts">
import { onMounted, ref } from "vue";
import layout from "../components/layout.vue";
import { useAuth0 } from "@auth0/auth0-vue";
import { getFiles } from "../services/files";
import type { File } from "../shared/types";
import FileGrid from "../components/file-grid.vue";

const { user, getAccessTokenSilently } = useAuth0();

const token = ref<string | null>(null);

const files = ref<File[]>([]);

onMounted(async () => {
	try {
		token.value = await getAccessTokenSilently();
		files.value = await getFiles();
	} catch (error) {
		console.error(error);
	}
});
</script>

<template>
	<layout>
		<div v-if="user">
			<FileGrid :files="files" />
			<p>
				{{ token }}
			</p>
			<p>
				{{ JSON.stringify(user, null, 1) }}
			</p>
		</div>
	</layout>
</template>
