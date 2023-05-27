<script>
	import { Router, Link, Route } from 'svelte-routing';
	import RSVPForm from './RSVPForm.svelte';
	import Attendees from './Attendees.svelte';
</script>
  
<style>
  @import url('https://fonts.googleapis.com/css2?family=Great+Vibes&family=Oxygen&display=swap');
	* {
		font-family: 'Oxygen', cursive;
	}
</style>
  

  <Router>
	<main>
	  <Route path="/" component="{RSVPForm}" />
	  <Route path="/attendees" component="{Attendees}" />
	</main>
  </Router>
  