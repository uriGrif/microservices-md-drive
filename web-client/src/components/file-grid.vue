<script setup lang="ts">
import { ref } from "vue";
import type { File } from "../shared/types";
import FileCard from "./file-card.vue";
import themeOverrides from "../shared/formThemeOverride";
import { NInput, NSelect, NConfigProvider } from "naive-ui";
import btn from "./buttons/btn.vue";

defineProps<{
	files: Array<File>;
}>();

const sortOptions = [
	{ label: "Name (asc)", value: "name-asc" },
	{ label: "Name (desc)", value: "name-desc" },
	{ label: "Created At (asc)", value: "createdAt-asc" },
	{ label: "Created At (desc)", value: "createdAt-desc" },
	{ label: "Updated At (asc)", value: "updatedAt-asc" },
	{ label: "Updated At (desc)", value: "updatedAt-desc" }
];
const filterOptions = [
	{ label: "All", value: "" },
	{ label: "My Files", value: "my-files" },
	{ label: "Shared with Me", value: "shared-with-me" }
];
const searchString = ref<string>("");
const sortBy = ref<string>("name-asc");
const filter = ref<string>("");
</script>

<template>
	<div class="file-grid-header">
		<div class="header-content">
			<h2>Home</h2>
			<btn text="New Document" />
		</div>
		<n-config-provider :theme-overrides="themeOverrides">
			<n-input
				v-model:value="searchString"
				placeholder="Search your documents"
			/>
			<div class="controls">
				<div class="control">
					<span style="width: 80px">Sort By:</span>
					<n-select
						class="sortby-select"
						v-model:value="sortBy"
						:options="sortOptions"
					/>
				</div>
				<div class="control">
					<span>Filter:</span>
					<n-select
						class="sortby-select"
						v-model:value="filter"
						:options="filterOptions"
					/>
				</div>
			</div>
		</n-config-provider>
	</div>
	<div class="file-grid">
		<FileCard
			v-for="file in files"
			:key="file.id"
			:id="file.id"
			:name="file.name"
			:createdAt="file.createdAt"
			:updatedAt="file.updatedAt"
			:owner="file.userId"
		/>
	</div>
</template>

<style scoped>
.file-grid {
	display: grid;
	grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
	gap: 1rem;
}

.file-grid-header {
	margin-bottom: 2rem;

	.controls {
		display: flex;
		flex-direction: row;
	}

	.control {
		display: flex;
		align-items: center;
		gap: 1rem;
		margin-top: 1rem;
		width: 250px;
		margin-right: 1rem;
	}
}

.header-content {
	display: flex;
	justify-content: space-between;
	align-items: center;
	background-color: var(--primary-dark-color);
	border-radius: var(--border-radius);
	padding: 1rem;
	margin-bottom: 1rem;

	h2 {
		margin: 0px;
	}
}
</style>
