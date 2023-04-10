<script>
	import { onMount } from 'svelte';
	let name = '';
	let attending = false;
	let message = '';
	let guests = 0;
    let restrictions = '';
  
	async function handleSubmit() {
		const response = await fetch(`/api/rsvp`, {
			method: 'POST',
			headers: {
			'Content-Type': 'application/json',
			},
			body: JSON.stringify({ name, attending, guests, restrictions }),
		});
  
	  if (response.ok) {
		if(attending) {
			message = `Thanks for your RSVP, ${name}! We're looking forward to seeing you!`;
		} else {
			message = `Thanks for your RSVP, ${name}. You'll be missed!`
		}
	  } else {
		message = 'Something went wrong. Please try again later.';
	  }
	}
</script>

  <main>
	
  <div class="container mt-5">
	<div class="row">
	  <div class="col-md-8 offset-md-2">
		<div class="card">
		  <div class="card-body">
			<h2 class="card-title text-center mb-4">RSVP</h2>
			<form on:submit|preventDefault={handleSubmit}>
			  <div class="mb-3">
				<label for="name" class="form-label">Name</label>
				<input
				  type="text"
				  class="form-control"
				  id="name"
				  bind:value={name}
				  on:input={() => { message = ""; }}
				  required
				/>
			  </div>
			  <div class="mb-3">
				<label for="guests" class="form-label"># of Guests</label>
				<input
				  type="number"
				  class="form-control"
				  id="guests"
				  bind:value={guests}
				  on:input={() => { message = ""; }}
				  required
				/>
			  </div>
              <div class="mb-3">
				<label for="restrictions" class="form-label">Dietary Restrictions</label>
				<input
				  type="text"
				  class="form-control"
				  id="restrictions"
				  bind:value={restrictions}
				  on:input={() => { message = ""; }}
				  required
				/>
			  </div>
			  <div class="mb-3 form-check">
				<input
				  type="checkbox"
				  class="form-check-input"
				  id="attending"
				  bind:checked={attending}
				/>
				<label class="form-check-label" for="attending">Attending</label>
			  </div>
			  <div class="d-grid gap-2">
				<button type="submit" class="btn btn-primary">Submit</button>
			  </div>
			</form>
			{#if message}
			  <div class="alert alert-{message.includes('Thanks') ? 'success' : 'danger'} mt-3" role="alert">
				{message}
			  </div>
			{/if}
		  </div>
		</div>
	  </div>
	</div>
  </div>
</main>
  
  
  