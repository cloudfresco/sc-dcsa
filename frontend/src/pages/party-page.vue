<template>
  <PageLayout>
    <div class="content-layout">
      <h1 id="page-title" class="content__title">Party Page</h1>
      <div class="content__body">
        <p id="page-description">
          <span
            >This page retrieves a <strong>Party message</strong> from an
            external API.</span
          >
          <span
            ><strong
              >Only authenticated users can access this page.</strong
            ></span
          >
        </p>
        <CodeSnippet title="Party Message" :code="message" />
      </div>
    </div>
  </PageLayout>
</template>

<script setup>
import CodeSnippet from "@/components/code-snippet.vue";
import PageLayout from "@/components/page-layout.vue";
import { getPartyResource } from "@/services/message.service";
import { ref } from "vue";
import { useAuth0 } from "@auth0/auth0-vue";

const message = ref("");

const getMessage = async () => {
  const { getAccessTokenSilently } = useAuth0();
  const accessToken = await getAccessTokenSilently();
  alert(accessToken)
  const { data, error } = await getPartyResource(accessToken);

  if (data) {
    message.value = JSON.stringify(data, null, 2);
  }

  if (error) {
    message.value = JSON.stringify(error, null, 2);
  }
};

getMessage();
</script>
