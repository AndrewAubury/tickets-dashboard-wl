<div class="wrapper">
  {#if imageUrl !== null}
    <img src={imageUrl} class="logo" bind:this={logo} on:error={useDefaultLogo}/>
  {:else}
    <img src="/assets/img/grey.png" class="logo"/>
  {/if}
  <div class="details">
    <div class="title-row">
      <span class="title">{name}</span>
      {#if builtIn}
        <Badge colour="#0c8f43">Built-In</Badge>
      {/if}
      {#if added}
        <Badge>Active</Badge>
      {/if}
      {#if guildCount !== undefined}
        <Badge>
          <div class="guild-count">
            <i class="fas fa-server"></i>
            {guildCount}
          </div>
        </Badge>
      {/if}
    </div>

    {#if showAuthor && author}
      <div class="author">
        <a href="https://discord.com/users/{author.id}" class="link" style="gap: 4px">
          <img src="https://cdn.discordapp.com/avatars/{author.id}/{author.avatar}.webp" class="author-avatar"
               alt="Author avatar" on:error={(e) => handleAvatarError(e, author.id)}/>
          <b>{author.global_name || author.username}</b>
        </a>
      </div>
    {:else if showAuthor}
      <div class="author">
        <a href="https://discord.com/users/{ownerId}" class="link" style="gap: 4px">
          <img src="https://cdn.discordapp.com/embed/avatars/0.png" class="author-avatar"
               alt="Author avatar" />
          <b>Unknown User</b>
        </a>
      </div>
    {/if}

    <span class="description">
      <slot name="description"></slot>
    </span>

    {#if !hideLinks}
      <div class="links">
        {#if builtIn}
          <a href="{viewLink}" target="_blank" class="link-blue">View</a>
        {:else if added}
          <Navigate to="/manage/{guildId}/integrations/view/{integrationId}" styles="link-blue">View</Navigate>
          <Navigate to="/manage/{guildId}/integrations/manage/{integrationId}" styles="link-blue">Configure</Navigate>
          <a href="#" class="link-blue" on:click={() => dispatch("remove", {})}>Remove</a>
        {:else}
          {#if owned}
            <Navigate to="/manage/{guildId}/integrations/view/{integrationId}" styles="link-blue">Preview</Navigate>
            <Navigate to="/manage/{guildId}/integrations/configure/{integrationId}" styles="link-blue">Configure
            </Navigate>
          {:else}
            <Navigate to="/manage/{guildId}/integrations/view/{integrationId}" styles="link-blue">View</Navigate>
          {/if}
          <Navigate to="/manage/{guildId}/integrations/activate/{integrationId}" styles="link-blue">Add to server
          </Navigate>
        {/if}
      </div>
    {/if}
  </div>
</div>

<script>
    import Badge from "../Badge.svelte";
    import {Navigate} from "svelte-router-spa";
    import {createEventDispatcher} from "svelte";
    import {getIconUrl} from "../../js/icons";

    const dispatch = createEventDispatcher();

    export let guildId;

    export let integrationId;
    export let name;
    export let imageUrl;
    export let ownerId;
    export let owned = false;
    export let guildCount;
    export let author;
    export let showAuthor = false;

    export let added = false;
    export let builtIn = false;
    export let hideLinks = false;
    export let viewLink;

    let logo;

    function useDefaultLogo() {
        logo.src = "/assets/img/grey.png";
    }

    function handleAvatarError(ev, id) {
        const src = getIconUrl(id, "");
        if (ev.target.src === src) { // Setting onerror to null does not work with svelte
            return;
        }

        ev.target.src = src;
    }
</script>

<style>
    .wrapper {
        display: flex;
        flex-direction: column;
        border-radius: 10px;

        background-color: #272727 !important;
        box-shadow: 0 4px 4px rgba(0, 0, 0, 0.25);
        transition: all .3s ease-in-out;

        height: 100%;
    }

    .logo {
        width: 90%;
        border-radius: 10px 10px 0 0;
        max-height: 150px;
        object-fit: cover;
        margin: auto;
    }

    .details {
        display: flex;
        flex-direction: column;
        padding: 10px 20px 10px 20px;
        height: 100%;
    }

    .title-row {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 10px;
    }

    .title {
        font-size: 18px;
        font-weight: bolder;
        word-break: break-all;
    }

    .description {
        flex: 1;
        font-size: 14px;
        color: #b9bbbe;
    }

    .guild-count {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 4px;
    }

    .links {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-around;
    }

    .author {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 4px;
    }

    .author-avatar {
        height: 24px;
        border-radius: 50%;
    }
</style>