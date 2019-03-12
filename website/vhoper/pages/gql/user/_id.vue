<template>
  <!-- Apollo watched Graphql query -->
  <ApolloQuery
    :query="require('../../../graphql/User.gql')"
    :variables="{
      id: $route.params.id
    }"
  >
    <template slot-scope="{ result: { loading, error, data } }">
      <!-- Loading -->
      <div v-if="loading" class="loading apollo">
        Loading...
      </div>

      <!-- Error -->
      <div v-else-if="error" class="error apollo">
        An error occured
      </div>

      <!-- Result -->
      <div v-else-if="data" class="result apollo">
        <div>{{ data.getUser.Name }}</div>
        <div>{{ data.getUser.ID }}</div>
        <div>{{ data.getUser.Phone }}</div>
      </div>

      <!-- No result -->
      <div v-else class="no-result apollo">
        No result :(
      </div>
    </template>
  </ApolloQuery>
</template>

<script>
export default {}
</script>

<style>
</style>
