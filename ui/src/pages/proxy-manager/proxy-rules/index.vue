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
                <FormItem label="type">
                    <Select v-model="formItem.select" multiple>
                        <Option value="ch.proxy.ru">ch.proxy.ru</Option>
                        <Option value="ru.proxy.ru">ru.proxy.ru</Option>
                        <Option value="en.proxy.ru">en.proxy.ru</Option>
                    </Select>
                </FormItem>
                <FormItem label="value">
                    <Input v-model="formItem.value" placeholder="Enter something..."></Input>
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
          type:'',
          value: '',
          policy: '',
          note: '',
        },
        columns: [
          {
            title: 'Match Type',
            slot: 'type'
          },
          {
            title: 'Value',
            key: 'value'
          },
          {
            title: 'Policy',
            key: 'policy'
          },
          {
            title: 'Note',
            key: 'note'
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
            type: 'DOMAIN',
            value: 'gs.apple.com',
            policy: "Proxy",
            note: ""
          },
          {
            type: 'DOMAIN',
            value: 'gs.apple.com',
            policy: "Proxy",
            note: ""
          },
          {
            type: 'DOMAIN',
            value: 'tx.jd.com',
            policy: "REJECT",
            note: ""
          },
          {
            type: 'DOMAIN-KEYWORD',
            value: 'nowtv100',
            policy: "Proxy",
            note: ""
          },
        ]
      }
    },
    methods: {
      ok () {
        this.data.push({
          type: this.formItem.type,
          value: this.formItem.value,
          policy:  this.formItem.policy,
          note: this.formItem.note,
        })
        this.$refs["formCustom"].resetFields()
        this.modal = false
      },
      cancel () {
        this.modal = false
      },
      show (index) {
        this.$Modal.info({
          title: 'Edit rule',
        })
      },
      remove (index) {
        this.data.splice(index, 1);
      }
    }
  }
</script>
