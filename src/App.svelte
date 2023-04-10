<script>
	import { Router, Link, Route } from 'svelte-routing';
	import RSVPForm from './RSVPForm.svelte';
	import Attendees from './Attendees.svelte';
  </script>
  
  <style>
	/* ... (same as before) */
  </style>
  
  <Router>
	<nav class="navbar navbar-expand-lg navbar-light bg-light">
	  <div class="container">
		<Link class="navbar-brand" to="/">RSVP</Link>
		<div class="collapse navbar-collapse">
		  <ul class="navbar-nav me-auto mb-2 mb-lg-0">
			<li class="nav-item">
			  <Link class="nav-link" to="/attendees">Attendees</Link>
			</li>
		  </ul>
		</div>
	  </div>
	</nav>
  
	<main>
	  <Route path="/" component="{RSVPForm}" />
	  <Route path="/attendees" component="{Attendees}" />
	</main>
  </Router>
  