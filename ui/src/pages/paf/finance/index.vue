<template>
    <Table :columns="columns1" :data="data1"></Table>
</template>
<script>
  export default {
    data () {
      return {
        columns1: [
          {
            title: 'Date',
            key: 'date'
          },
          {
            title: 'Amont',
            key: 'amount'
          },
          {
            title: 'Type',
            key: 'type'
          }
        ],
        data1: [
          {
            type: 'Pay Extra GB',
            amount: 18,
            date: '2021-10-03'
          },
          {
            type: 'Pay Extra GB',
            amount: 24,
            date: '2021-10-01'
          },
          {
            type: 'Pay Extra GB',
            amount: 30,
            date: '2021-10-02'
          },
          {
            type: 'Pay Extra GB',
            amount: 26,
            date: '2012-10-04'
          }
        ]
      }
    }
  }
</script>
