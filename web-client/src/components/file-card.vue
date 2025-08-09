<script setup lang="ts">
import {
	NIcon,
	NConfigProvider,
	type GlobalThemeOverrides,
	NDropdown
} from "naive-ui";
import { Document, Calendar, Time, OverflowMenuVertical } from "@vicons/carbon";

defineProps<{
	id: string;
	name: string;
	createdAt: string;
	updatedAt: string;
	owner: string;
}>();

function formatDate(dateString: string) {
	const date = new Date(dateString);
	return date.toLocaleDateString("en-US", {
		month: "short",
		day: "numeric",
		year: "numeric"
	});
}

const dropdownTheme: GlobalThemeOverrides = {
	Dropdown: {
		color: "var(--primary-light-color)", // Background color
		textColor: "var(--primary-dark-color)", // Text color
		optionColorHover: "var(--text-color)", // Hover background color
		optionTextColorHover: "var(--primary-dark-color)", // Hover background color
		optionTextColor: "var(--primary-dark-color)" // Option text color
	}
};

const handleSelect = (key: string | number) => {
	switch (key) {
		case "edit":
			break;
		case "share":
			break;
		case "delete":
			break;
	}
};

const dropdownOptions = [
	{
		label: "Edit",
		key: "edit"
	},
	{
		label: "Share",
		key: "share"
	},
	{
		label: "Delete",
		key: "delete"
	}
];
</script>

<template>
	<div className="file-card">
		<div className="file-card-header">
			<div className="file-header-content">
				<div className="file-title-section">
					<n-icon class="meta-icon">
						<Document />
					</n-icon>
					<h3 className="file-name">{{ name }}</h3>
				</div>
			</div>
		</div>
		<div className="file-card-content">
			<div className="file-meta">
				<div className="file-meta-item">
					<n-icon class="meta-icon">
						<Calendar />
					</n-icon>
					<span>Created {{ formatDate(createdAt) }}</span>
				</div>
				<div className="file-meta-item">
					<n-icon class="meta-icon">
						<Time />
					</n-icon>
					<span>Updated {{ formatDate(updatedAt) }}</span>
				</div>
			</div>
		</div>
		<n-config-provider :theme-overrides="dropdownTheme">
			<n-dropdown
				trigger="click"
				:options="dropdownOptions"
				@select="handleSelect"
				placement="right-start"
			>
				<n-icon class="menu-icon">
					<OverflowMenuVertical />
				</n-icon>
			</n-dropdown>
		</n-config-provider>
	</div>
</template>

<style scoped>
.file-card {
	background: var(--primary-dark-color);
	border: 1px solid var(--primary-dark-color);
	border-radius: var(--border-radius);
	cursor: pointer;
	transition: all 0.2s;
	position: relative;
}

.file-card:hover {
	border-color: color-mix(in srgb, var(--accent-color) 50%, transparent);
	box-shadow: 0 10px 15px -3px color-mix(in srgb, var(--accent-color) 10%, transparent);
}

.file-card-header {
	padding: 1rem 1rem 0.75rem 1rem;
}

.file-header-content {
	display: flex;
	align-items: flex-start;
	justify-content: space-between;
}

.file-title-section {
	display: flex;
	align-items: center;
	gap: 0.75rem;
	min-width: 0;
}

.file-icon {
	color: var(--accent-color);
	width: 1.25rem;
	height: 1.25rem;
	flex-shrink: 0;
}

.file-name {
	font-weight: 500;
	font-size: 0.875rem;
	color: var(--text-color);
	transition: color 0.2s;
	line-height: 1.2;
	display: -webkit-box;
	-webkit-box-orient: vertical;
	overflow: hidden;
}

.file-card:hover .file-name {
	color: var(--accent-color);
}

.file-card-content {
	padding: 0 1rem 1rem 1rem;
}

.file-meta {
	display: flex;
	flex-direction: column;
	gap: 0.75rem;
}

.file-meta-item {
	display: flex;
	align-items: center;
	gap: 0.5rem;
	font-size: 0.75rem;
	color: var(--primary-light-color);
}

.menu-icon {
	position: absolute;
	top: 0.5rem;
	right: 0.5rem;
	color: var(--primary-light-color);
	cursor: pointer;
	font-size: 1.25rem;
}
</style>
