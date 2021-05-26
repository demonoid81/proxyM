<template>
    <Table :columns="columns" :data="data"></Table>
</template>

<script>
  export default {
    data () {
      return {
        columns: [
          {
            title: 'Proxy Addres',
            key: 'proxy'
          },
          {
            title: 'Traffic usage',
            key: 'traffic',
            children: [
              {
                title: 'Download',
                key: 'dowload'
              },
              {
                title: 'Upload',
                key: 'upload'
              },
              {
                title: 'Total',
                key: 'total'
              }
            ]
          },
        ],
        data: []
      }
    },
    mounted () {
      const min = 1024
      const max = 102400
      const data = [];
      for (let i = 0; i < 20; i++) {
        let dowload = Math.floor(Math.random() * (max - min + 1)) + min;
        let upload = Math.floor(Math.random() * (max - min + 1)) + min;
        data.push({
          key: i,
          proxy: 'ch.proxy.ru',
          dowload: dowload,
          upload: upload,
          total: dowload + upload
        });
      }
      this.data = data;
    }
  }
</script>
