<script lang="ts">
	import { CrudModes } from '$lib/models/crud_modes';
	import { Button, Navbar, NavUl } from 'flowbite-svelte';
	import { closeModal, Modals, openModal } from 'svelte-modals';
	import SubjectModal from './modals/SubjectModal.svelte';

	function showCreateSubjectModal() {
		openModal(SubjectModal, {
			mode: CrudModes.CREATE,
		});
	}
	// function showCreatePredicateModal() {
	// 	openModal(PredicateModal, {
	// 		mode: CrudModes.CREATE,
	// 	});
	// }
	// function showCreatePremiseModal() {
	// 	openModal(PremiseModal, {
	// 		mode: CrudModes.CREATE,
	// 	});
	// }
	// function showCreatePropositionModal() {
	// 	openModal(PropositionModal, {
	// 		mode: CrudModes.CREATE,
	// 	});
	// }
	// function showCreateConditionalStatementModal() {
	// 	openModal(ConditionalStatementModal, {
	// 		mode: CrudModes.CREATE,
	// 	});
	// }
	// function showCreateArumentModal() {
	// 	openModal(ArumentModal, {
	// 		mode: CrudModes.CREATE,
	// 	});
	// }

	const actionItems = [
		{ text: 'Create Subject', handleClick: showCreateSubjectModal },
		{
			text: 'Create Predicate',
			handleClick: function () {
				alert('wow');
			}
		},
		{
			text: 'Create Premise',
			handleClick: function () {
				alert('wow');
			}
		},
		{
			text: 'Create Proposition',
			handleClick: function () {
				alert('wow');
			}
		},
		{
			text: 'Create Conditional Statement',
			handleClick: function () {
				alert('wow');
			}
		},
		{
			text: 'Create Argument',
			handleClick: function () {
				alert('wow');
			}
		}
	];
</script>

<Navbar>
	<NavUl>
		{#each actionItems as actionItem}
			<Button on:click={actionItem.handleClick}>{actionItem.text}</Button>
		{/each}
	</NavUl>
</Navbar>

<Modals>
	<div slot="backdrop" class="backdrop" on:click={closeModal} />
</Modals>

<style>
	.backdrop {
	  position: fixed;
	  top: 0;
	  bottom: 0;
	  right: 0;
	  left: 0;
	  background: rgba(0,0,0,0.50)
	}
  </style>