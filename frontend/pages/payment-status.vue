<template>
  <div class="min-h-screen flex items-center justify-center">
    <div class="bg-green-100 p-8 rounded-md shadow-md">
      <h1 class="text-2xl font-bold text-green-700">Payment Successful!</h1>
      <p class="mt-4">
        Thank you for purchasing a ticket for our event . Your transaction
        reference is: {{ txRef }}
      </p>
      <!-- <p class="mt-2">An email confirmation has been sent to {{ userEmail }}.</p> -->
      <router-link
        to="/"
        class="mt-4 inline-block bg-green-700 text-white px-4 py-2 rounded-md hover:bg-green-800"
      >
        Back to Home
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from "vue";
import { useRoute } from "vue-router";
import RESERVE_TICKET from "~/graphql/mutations/ReserveTicket.gql";
import GetSingleTransaction from "~/graphql/query/GetSingleTransaction.gql";
import { useQuery, useMutation } from "@vue/apollo-composable";
import { useAuthStore } from "~/store";

const txRef = ref("");
const event_id = ref("");
const total_price = ref(0);
const quantity = ref(0);
const route = useRoute();
const user = useAuthStore();
const currentUser = user.id;
const currentUserRole = user.role;

const { mutate: reserveTicket } = useMutation(
  RESERVE_TICKET,
  {},
  {
    context: {
      headers: {
        "x-hasura-user-id": currentUser,
        "x-hasura-role": currentUserRole,
        Authorization: `Bearer ${user.token}`,
      },
    },
  }
);

onMounted(() => {
  txRef.value = route.query.tx_ref || "No transaction reference found";
  handlePaymentStatus();
});

const handlePaymentStatus = async () => {
  const params = new URLSearchParams(window.location.search);
  const status = params.get("status");
  const tx_ref = params.get("tx_ref");
  if (!tx_ref) {
    console.error("Transaction reference not found.");
    return;
  }
  try {
    const { onResult } = useQuery(
      GetSingleTransaction,
      {
        tx_ref,
      },
      {
        context: {
          headers: {
            "x-hasura-user-id": currentUser,
            "x-hasura-role": currentUserRole,
            Authorization: `Bearer ${user.token}`,
          },
        },
      }
    );
    onResult((result) => {
      if (result.data) {
        const data = result.data.transactions_by_pk;
        event_id.value = data.event_id;
        total_price.value = data.total_price;
        quantity.value = data.quantity;
        handleReserveTicket();
      }
    });
  } catch (error) {
    console.error("Failed to reserve the ticket:", error);
    toast.error("Failed to reserve ticket.");
  }
};

const handleReserveTicket = async () => {
  try {
    await reserveTicket(
      {
        event_id: event_id.value,
        quantity: quantity.value,
        total_price: total_price.value,
      },
      {
        context: {
          headers: {
            "x-hasura-user-id": currentUser,
            "x-hasura-role": currentUserRole,
            Authorization: `Bearer ${user.token}`,
          },
        },
      }
    );
    toast.success("Reservation completed successfully!");
  } catch (error) {
    console.error("Failed to reserve the ticket:", error);
    toast.error("Failed to reserve ticket.");
  }
};
</script>
