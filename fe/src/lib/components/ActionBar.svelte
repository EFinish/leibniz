<script lang="ts">
	import { Button, Navbar, NavLi, NavUl } from 'flowbite-svelte';

	const actionItems = [
		{text: 'Create Subject'},
		{text: 'Create Predicate'},
		{text: 'Create Premise'},
		{text: 'Create Proposition'},
		{text: 'Create Conditional Statement'},
		{text: 'Create Argument'}
	];
</script>

<Navbar>
	<NavUl>
		{#each actionItems as actionItem}
			<Button>{actionItem.text}</Button>
		{/each}
	</NavUl>
</Navbar>
