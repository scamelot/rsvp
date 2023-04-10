<script>
    import { onMount } from 'svelte';
  
    let attendees = [];
    let nonAttendees = [];
  
    onMount(async () => {
      const response = await fetch('/api/rsvp');
      const data = await response.json();
      attendees = data.filter(person => person.attending);
      nonAttendees = data.filter(person => !person.attending);
    });
  </script>
  
  <div class="container mt-5">
    <h2 class="mb-4">Attendees</h2>
    <table class="table table-striped">
        <thead>
          <tr>
            <th>Name</th>
            <th>Number of Guests</th>
            <th>Restrictions</th>
          </tr>
        </thead>
        <tbody>
          {#each attendees as attendee}
            <tr>
              <td>{attendee.name}</td>
              <td>{attendee.guests}</td>
              <td>{attendee.restrictions}</td>
            </tr>
          {/each}
        </tbody>
      </table>
  
    <h2 class="mb-4">Non-Attendees</h2>
    <ul class="list-group">
      {#each nonAttendees as nonAttendee}
        <li class="list-group-item">{nonAttendee.name}</li>
      {/each}
    </ul>
  </div>
  