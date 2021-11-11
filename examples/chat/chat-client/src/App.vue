<script setup lang="ts">
import { ChatClient } from "../gen/chat.client";
import { HrpcTransport } from "@harmony-dev/transport-hrpc";
import { onMounted, reactive, ref } from "vue";
const client = new ChatClient(
  new HrpcTransport({
    baseUrl: "http://127.0.0.1:2289",
    insecure: true
  })
);

const content = ref("");
const msgs = reactive<string[]>([]);

onMounted(() => {
  client.streamMessages({}).responses.onMessage((msg) => {
    msgs.push(msg.content);
  });
});

const send = (ev: KeyboardEvent) => {
  if (ev.key !== "Enter") return;
  client.sendMessage({
    content: content.value,
  });
  content.value = "";
};
</script>

<template>
  <div class="h-100vh w-100vw bg-surface-900 flex flex-col justify-center p-3">
    <div class="flex-1 p-3 flex flex-col gap-2 overflow-auto">
      <p class="p-3 max-w-30ch rounded-md bg-surface-800" v-for="m in msgs" :key="m">{{ m }}</p>
    </div>
    <input
      class="
        p-2
        bg-surface-700
        rounded-md
        focus:outline-none focus:ring-3
        ring-secondary-400
	mt-2
      "
      v-model="content"
      @keydown="send"
    />
  </div>
</template>
