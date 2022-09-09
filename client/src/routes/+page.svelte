<script>
  const URL = 'http://127.0.0.1:5050/';

  let home = fetch(URL).then(
    res => res.json()
  )
</script>

<svelte:head>
  <title>Bot</title>
</svelte:head>

{#await home}
  <p>Loading home</p>
{:then data} 
  <h1>{data.message}</h1>
{/await}