<script setup lang="ts">
import { ref } from "vue";
import type { CreateFileDTO, EditFileDTO, File } from "../shared/types";
import FileCard from "./file-card.vue";
import { NInput, NSelect, useMessage } from "naive-ui";
import btn from "./buttons/btn.vue";
import { createFile, editFile, deleteFile } from "../services/files";
import createFileModal from "./modals/create-file-modal.vue";
import editFileModal from "./modals/edit-file-modal.vue";
import deleteFileModal from "./modals/delete-file-modal.vue";

const emit = defineEmits(["refreshFiles"]);

const createFileModalRef = ref(
	<InstanceType<typeof createFileModal> | null>null
);
const editFileModalRef = ref(<InstanceType<typeof editFileModal> | null>null);
const deleteFileModalRef = ref(
	<InstanceType<typeof deleteFileModal> | null>null
);
const message = useMessage();

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

const create = async (file: CreateFileDTO) => {
	try {
		await createFile(file);
		emit("refreshFiles");
		createFileModalRef.value?.close();
	} catch (error) {
		console.log(error);
		message.error("Failed to create file");
	}
};

const openEditModal = (file: File) => {
	editFileModalRef.value?.open(file);
};

const edit = async (id: string, file: EditFileDTO) => {
	try {
		await editFile(id, file);
		emit("refreshFiles");
		editFileModalRef.value?.close();
	} catch (error) {
		console.log(error);
		message.error("Failed to edit file");
	}
};

const openDeleteModal = (file: File) => {
	deleteFileModalRef.value?.open(file);
};

const deleteF = async (id: string) => {
	try {
		await deleteFile(id);
		emit("refreshFiles");
		deleteFileModalRef.value?.close();
	} catch (error) {
		console.log(error);
		message.error("Failed to delete file");
	}
};
</script>

<template>
	<div class="file-grid-header">
		<div class="header-content">
			<h2>Home</h2>
			<btn text="New File" @click="createFileModalRef?.open()" />
		</div>
		<n-input v-model:value="searchString" placeholder="Search your files" />
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
	</div>
	<div class="file-grid">
		<FileCard
			v-for="file in files"
			:key="file.id"
			:id="file.id"
			:name="file.name"
			:createdAt="file.created_at"
			:updatedAt="file.updated_at"
			:owner="file.userId"
			@editFile="openEditModal(file)"
			@deleteFile="openDeleteModal(file)"
		/>
	</div>
	<createFileModal ref="createFileModalRef" @createFile="create" />
	<editFileModal ref="editFileModalRef" @editFile="edit" />
	<deleteFileModal ref="deleteFileModalRef" @deleteFile="deleteF" />
</template>

<style scoped>
.file-grid {
	display: grid;
	grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
	gap: 1rem;
	max-height: 50vh;
	overflow-y: scroll;
	padding: 20px 0px;
	padding-right: 3px; /* for scrollbar */
	mask-image: linear-gradient(
		to bottom,
		transparent 0%,
		var(--background-color) 10px,
		var(--background-color) calc(100% - 10px),
		transparent 100%
	);
}

.file-grid::-webkit-scrollbar {
	width: 3px;
}

.file-grid::-webkit-scrollbar-track {
	background: transparent;
	border-radius: 10px;
}

.file-grid::-webkit-scrollbar-thumb {
	background: var(--primary-dark-color);
	border-radius: 10px;
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
