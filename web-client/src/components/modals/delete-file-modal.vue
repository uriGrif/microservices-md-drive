<script setup lang="ts">
import { ref } from "vue";
import { NModal } from "naive-ui";
import type { File } from "../../shared/types";
import btn from "../buttons/btn.vue";

const show = ref<boolean>(false);

const emit = defineEmits(["deleteFile"]);

const fileId = ref<string>("");
const fileName = ref<string>("");

const open = (file: File) => {
	show.value = true;
	fileId.value = file.id;
	fileName.value = file.name;
};

const close = () => {
	show.value = false;
};

const deleteFile = () => {
	emit("deleteFile", fileId.value);
};

defineExpose({ open, close });
</script>

<template>
	<n-modal
		v-model:show="show"
		title="Delete File"
		:bordered="false"
		preset="card"
		style="max-width: 600px"
	>
		<p>
			Are you sure you want to delete the file
			<span style="font-weight: 800">{{ fileName }}</span>
			? This action cannot be undone.
		</p>
		<btn
			:onClick="close"
			text="Cancel"
			type="primary-light"
			style="margin-top: 20px"
		/>
		<btn
			:onClick="deleteFile"
			text="Delete"
			type="secondary"
			style="margin-top: 20px"
		/>
	</n-modal>
</template>
