<script>
	import { onMount } from 'svelte';
	import { createEventDispatcher } from 'svelte';
	// import Background from '../public/flowers.svg';
	let name = '';
	let attending = true;
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
	  openModal();
	}

  let showModal = false;
  const dispatch = createEventDispatcher();

  function openModal() {
    showModal = true;
  }

  function closeModal() {
	message = '';
    showModal = false;
  }

</script>
<style>
.modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 999;
  }

  .modal-content {
    background-color: white;
    padding: 20px;
    border-radius: 4px;
    width: 400px;
    max-width: 90%;
  }
	.image {
	display: block; /* Set display property to block or inline-block to ensure visibility on non-mobile devices */
	}

	@media (max-width: 767px) { /* Adjust the max-width value as needed for your desired mobile screen size */
	.image {
		display: none; /* Hide the image on mobile devices */
	}
	}
    @font-face {
		font-family: 'Bodoni Cyrillic';
		font-style: normal;
		font-weight: normal;
		src: local('Bodoni Cyrillic'), url('/BodoniCyrillic.ttf') format('truetype');
	}
		* {
          font-family: 'Bodoni Cyrillic', serif;
      }
	@font-face {
		font-family: "TAN-PEARL";
		font-style: normal;
		font-weight: normal;
		src: local('TAN-PEARL'), url('/TAN-PEARL.ttf') format('truetype');
	}
	  main {
		  background-color: #fdfefe;
		  color: #466f21;
		  font-size: 1.5rem;
		  opacity: 0.9;
		  width: 101%;
		  height: 120vh;
	  }
	  .background {
		position: fixed;
		top: 0;
		right: 0;
		width: 50%; /* Adjust the width as needed */
		height: 50%; /* Adjust the height as needed */
		z-index: -1; /* To move the SVG behind other elements */
	  } 
  </style>
  <div class="background">
	<img src="flowers.png" class="image" alt="sunflowers" />
  </div>	
  <main>
  <div class="container">
	<div class="row">
	  <div class="col-md-8 offset-md-2">
		<div class="d-flex justify-content-center"><img src="/engagementinvite.png" alt="Steve and Rachel's Engagement Party" width="125%" height="auto" />
		</div>
		  <div class="card-body">
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
				/>
			  </div>
			  <div class="mb-3 form-check form-switch form-switch-xl">
				<input
				  type="checkbox"
				  class="form-check-input"
				  id="attending"
				  bind:checked={attending}
				/>
				<label class="form-check-label" for="attending">Attending</label>
			  </div>
			  <div class="d-grid gap-2">
				<button type="submit" class="btn btn-{attending ? 'success' : 'danger'} ">
                    {#if attending}
                        Send your RSVP
                    {:else}
                        Gracefully Decline
                    {/if}
                        </button>
			  </div>
			</form>
			{#if message}
			  <div class="alert alert-{message.includes('Thanks') ? 'success' : 'danger'} mt-3" role="alert">
				<div class="modal">
					<div class="modal-content">
					  <p>{message}</p>
					  <button on:click={closeModal}>Close</button>
					</div>
				  </div>
			  </div>
			{/if}

		</div>
	  </div>
	</div>
  </div>
</main>
  
  
  
