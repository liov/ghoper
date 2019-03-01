<template>
  <!-- Apollo watched Graphql query -->
  <ApolloQuery
    :query="require('../../graphql/HelloWorld.gql')"
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
        {{ data.getUser.Name }}
        <div>{{ data.getUser.ID }}</div>
      </div>

      <!-- No result -->
      <div v-else class="no-result apollo">
        No result :(
      </div>
    </template>
  </ApolloQuery>
</template>

<script>
import ApolloExample from '../../components/ApolloExample.vue'

export default {
  components: {
    ApolloExample
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
