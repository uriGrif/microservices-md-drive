<script setup lang="ts">
import { ref } from "vue";
import {
	NModal,
	NForm,
	NFormItem,
	NInput,
	type FormRules,
	type FormInst,
	type FormItemRule
} from "naive-ui";
import btn from "../../components/buttons/btn.vue";
import type { EditFileDTO, File } from "../../shared/types";
import modalThemes from "./modals-themes";

const show = ref<boolean>(false);
const formRef = ref<FormInst | null>(null);

const emit = defineEmits(["editFile"]);

const file = ref<EditFileDTO>({
	name: ""
});
const fileId = ref<string>("");

const open = (originalFile: File) => {
	show.value = true;
	fileId.value = originalFile.id;
	file.value = {
		name: originalFile.name
	};
};

const close = () => {
	show.value = false;
	file.value = {
		name: ""
	};
};

defineExpose({ open, close });

const rules: FormRules = {
	name: {
		trigger: ["input", "blur"],
		type: "string",
		validator(_rule: FormItemRule, value: string) {
			if (value.length === 0) {
				return new Error("You must provide a file name");
			}
			return true;
		}
	}
};

const editFile = () => {
	formRef.value?.validate(async errors => {
		if (!errors) emit("editFile", fileId.value, file.value);
	});
};
</script>

<template>
	<n-modal
		v-model:show="show"
		title="Edit File"
		:bordered="false"
		preset="card"
		style="max-width: 600px"
	>
		<n-form :model="file" :rules="rules" ref="formRef">
			<n-form-item path="name">
				<n-input
					v-model:value="file.name"
					placeholder="File Name"
					:themeOverrides="modalThemes.Input"
				/>
			</n-form-item>
		</n-form>
		<btn
			:onClick="editFile"
			text="Edit"
			type="secondary"
			style="margin-top: 20px"
		/>
	</n-modal>
</template>
