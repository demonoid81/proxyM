<template>
    <div style="display: flex; flex-direction: column">
        <Button type="primary" style="width: 150px" @click="modal = true">Add IP address</Button>
        <Modal
                v-model="modal"
                title="Add IP address"
                ok-text="OK"
                cancel-text="Cancel"
                @on-ok="ok"
                @on-cancel="cancel">
            <Form ref="formCustom" :model="formItem" :label-width="100">
                <FormItem label="Ip address">
                    <Input v-model="formItem.name" placeholder="Enter something..."></Input>
                </FormItem>
                <FormItem label="Proxies">
                    <Select v-model="formItem.select" multiple>
                        <Option value="ch.proxy.ru">ch.proxy.ru</Option>
                        <Option value="ru.proxy.ru">ru.proxy.ru</Option>
                        <Option value="en.proxy.ru">en.proxy.ru</Option>
                    </Select>
                </FormItem>
                <FormItem label="rule">
                    <Select v-model="formItem.rule">
                        <Option value="default">default</Option>
                        <Option value="only russian">only russian</Option>
                        <Option value="only ch domains">only ch domains</Option>
                    </Select>
                </FormItem>
            </Form>
        </Modal>
        <div style="height: 10px"></div>
        <Table border :columns="columns" :data="data">
            <template slot-scope="{ row }" slot="name">
                <strong>{{ row.name }}</strong>
            </template>
            <template slot-scope="{ row, index }" slot="action">
                <Button type="primary" size="small" style="margin-right: 5px" @click="show(index)">View</Button>
                <Button type="error" size="small" @click="remove(index)">Delete</Button>
            </template>
        </Table>
    </div>
</template>
<script>
  export default {
    data () {
      return {
        modal: false,
        formItem: {
          name: '',
          select: [],
          rule: '',
        },
        columns: [
          {
            title: 'IP',
            slot: 'name'
          },
          {
            title: 'Proxy Addresses',
            key: 'addresses'
          },
          {
            title: 'Rule',
            key: 'rule'
          },
          {
            title: 'Action',
            slot: 'action',
            width: 150,
            align: 'center'
          }
        ],
        data: [
          {
            name: '141.247.215.177',
            addresses: 'ch.proxy.ru',
            rule: "default"
          },
          {
            name: '198.88.196.9',
            addresses: 'ch.proxy.ru, ru.proxy.ru, en.proxy.ru',
            rule: "default"
          },
          {
            name: '167.114.7.157',
            addresses: 'ru.proxy.ru',
            rule: "default"
          },
          {
            name: '57.145.190.108',
            addresses: 'zh.proxy.ru',
            rule: "default"
          }
        ]
      }
    },
    methods: {
      ok () {
        let add = ""
        for (let item of this.formItem.select) {
          add = add + " " + item
        }
        this.data.push({
          name: this.formItem.name,
          addresses: add,
          rule:  this.formItem.rule
        })
        this.$refs["formCustom"].resetFields()
        this.modal = false
      },
      cancel () {
        this.modal = false
      },
      show (index) {
        this.$Modal.info({
          title: 'Password',
          content: 'Password：123456',
          okText: "Close"
        })
      },
      remove (index) {
        this.data.splice(index, 1);
      }
    }
  }
</script>
