<script lang="ts">
	import { CrudModes } from '$lib/models/crud_modes';
	import { closeModal } from 'svelte-modals';
	import type { Subject } from '../../proto/argumentaccess/v1/argument';
	import axios from "axios";


	export let isOpen: boolean;
	export let subject: Subject = { id: '', body: '' };
	export let mode: CrudModes = CrudModes.READ;

	function createSubject () {}
</script>

{#if isOpen}
	<div role="dialog" class="modal">
		<div class="contents">
			{#if mode === CrudModes.CREATE}
				<h2>Create Subject</h2>
				<div>
					<label for="subject-body">Body</label>
					<input bind:value={subject.body} />
				</div>
				<div class="actions">
					<button on:click={createSubject}>Create</button>
					<button on:click={closeModal}>Cancel</button>
				</div>
			{:else if mode === CrudModes.READ}
				<h2>About Subject</h2>
				<div class="actions">
					<button on:click={closeModal}>OK</button>
				</div>
			{:else if mode === CrudModes.UPDATE}
				<h2>Edit Subject</h2>
				<div class="actions">
					<button>Update</button>
					<button on:click={closeModal}>Cancel</button>
				</div>
			{:else if mode === CrudModes.DELETE}
				<h2>Delete Subject</h2>
				<div class="actions">
					<button>Delete</button>
					<button on:click={closeModal}>Cancel</button>
				</div>
			{/if}
		</div>
	</div>
{/if}

<style>
	input {
		border: 1px solid black;
	}

	.modal {
		position: fixed;
		top: 0;
		bottom: 0;
		right: 0;
		left: 0;
		display: flex;
		justify-content: center;
		align-items: center;

		/* allow click-through to backdrop */
		pointer-events: none;
	}

	.contents {
		min-width: 240px;
		border-radius: 6px;
		padding: 16px;
		background: white;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		pointer-events: auto;
	}

	h2 {
		text-align: center;
		font-size: 24px;
	}

	p {
		text-align: center;
		margin-top: 16px;
	}

	.actions {
		margin-top: 32px;
		display: flex;
		justify-content: space-between;
		gap: 8px;
	}
</style>
