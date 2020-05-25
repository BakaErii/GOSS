<template>
	<div>
		<div class="layout">
			<Layout>
				<Content :style="{padding: '0 50px'}">
					<div>
						<Breadcrumb :style="{margin: '16px 0'}">
							<BreadcrumbItem>Status</BreadcrumbItem>
						</Breadcrumb>
						<div class="div-status">
							<Row :gutter="16" type="flex" justify="start">
								<i-col span="24" v-for="server in this.serverStatus" :key="server.name" style="margin-bottom: 15px;">
									<Card>
										<Row type="flex" justify="center" align="middle" :style="{margin: '10px 0'}">
											<i-col span="4">
												<h3>
													<Icon type="logo-buffer" :size="32" v-bind:color="nowTime-server.lastReportTime*1000>timeLimit*1000?'#ed4014':'#19be6b'"></Icon>{{ server.name }}
												</h3>
											</i-col>
											<i-col span="20">
												<p>
													{{ server.description }}
												</p>
											</i-col>
										</Row>
										<Row type="flex" justify="start" align="middle" :style="{margin: '15px 0'}">
											<i-col span="6">
												<center>
													<p>IP Address</p>
													<Tag color="default" size="medium">{{ server.ip }}</Tag>
												</center>
											</i-col>
											<i-col span="6">
												<center>
													<p>CPU Load</p>
													<Tag color="default" size="medium">{{ server.cpuUsage }}%</Tag>
												</center>
											</i-col>
											<i-col span="6">
												<center>
													<p>Memory Usage</p>
													<Tag color="default" size="medium">{{ server.memUsage>=1024?(server.memUsage/1024).toFixed(2):server.memUsage }}
														/ {{ server.memUsage>=1024?(server.maxMemLimit/1024).toFixed(2):server.maxMemLimit }}
														{{ server.memUsage>=1024?"GB":"MB" }}</Tag>
												</center>
											</i-col>
											<i-col span="6">
												<center>
													<p>Storage Usage</p>
													<Tag color="default" size="medium">{{ server.diskUsage>=1024?(server.diskUsage/1024).toFixed(2):server.diskUsage }}
														/ {{ server.diskUsage>=1024?(server.maxDiskLimit/1024).toFixed(2):server.maxDiskLimit }}
														{{ server.diskUsage>=1024?"GB":"MB" }}</Tag>
												</center>
											</i-col>
										</Row>
										<Row type="flex" justify="start" align="middle" :style="{margin: '15px 0'}">
											<i-col span="6">
												<center>
													<p>Network I/O</p>
													<Tag color="default" size="medium">{{ server.inBound.toFixed(2) }} / {{ server.outBound.toFixed(2) }} Mbps</Tag>
												</center>
											</i-col>
											<i-col span="6">
												<center>
													<p>Bandwidth Usage</p>
													<Tag color="default" size="medium">{{ server.inBoundTotalUsage>=1024?(server.inBoundTotalUsage/1024).toFixed(2):(server.inBoundTotalUsage).toFixed(2) }}
														/ {{ server.inBoundTotalUsage>=1024?(server.outBoundTotalUsage/1024).toFixed(2):(server.outBoundTotalUsage).toFixed(2) }}
														{{ server.inBoundTotalUsage>=1024?"GB":"MB" }}</Tag>
												</center>
											</i-col>
											<i-col span="12">
												<center>
													<p>Service Status</p>
													<Tag type="dot" v-for="(status,name) in server.serviceStatus" v-bind:key="name" v-bind:color="status==true?nowTime-server.lastReportTime*1000>30*1000?'error':'success':'error'">{{name}}</Tag>
												</center>
											</i-col>
										</Row>
										<Divider />
										<Row type="flex" justify="center" align="middle">
											<i-col span="24">
												Update Time <Time :time="server.lastReportTime * 1000" type="datetime" /> (<Time v-bind:time="server.lastReportTime"
												 :interval="1" />)
											</i-col>
										</Row>
									</Card>
								</i-col>
							</Row>
							<Spin size="large" fix v-if="spinShow"></Spin>
						</div>
					</div>
				</Content>
				<Footer class="layout-footer-center">Server Status</Footer>
			</Layout>
		</div>
	</div>
</template>

<script>
	export default {
		data() {
			return {
				serverStatus: null,
				spinShow: true,
				msgNetworkError: false,
				msgNetworkSuccess: false,
				lastReportError: false,
				timer: null,
				nowTime: null,
				timeLimit: window.g.timeLimit,
			}
		},
		mounted: function() {
			this.getServerStatus()
			this.timer = setInterval(this.getServerStatus, 1000 * window.g.timeDuration);
		},
		beforeDestroy() {
			clearInterval(this.timer)
		},
		methods: {
			getServerStatus() {
				var vm = this;
				this.$axios.post(window.g.apiUrl + "/status", {})
					.then(function(response) {
						vm.serverStatus = response.data['status']
						vm.spinShow = false
						if (vm.lastReportError == true) {
							vm.lastReportError = false;
							vm.showMsgNetworkSuccess();
						}
					})
					.catch(function(error) {
						vm.$Message.error(error.toString());
						vm.showMsgNetworkError();
						vm.lastReportError = true;
					});
				this.nowTime = new Date().getTime()
				return;
			},
			showMsgNetworkError() {
				var vm = this
				if (this.msgNetworkError == false) {
					this.msgNetworkError = true
					this.$Notice.error({
						title: 'Network error',
						desc: "Can't connect to server, retry in next duration.",
						top: 50,
						duration: 0,
						onClose: function() {
							vm.msgNetworkError = false
						}
					});
				}
			},
			showMsgNetworkSuccess() {
				var vm = this
				if (this.msgNetworkSuccess == false) {
					this.msgNetworkSuccess = true
					this.$Notice.success({
						title: 'Retry success',
						desc: 'Connect to server success.',
						duration: 0,
						onClose: function() {
							vm.msgNetworkSuccess = false
						}
					});
				}
			},
			showUnixTime(curTime) {
				var timeStamp = curTime * 1000
				var currTime = new Date(timeStamp)
				return currTime.toLocaleString()
			}
		}
	}
</script>

<style>
	.div-status {
		position: relative;
	}

	.layout {
		border: 1px solid #d7dde4;
		background: #f5f7f9;
		position: relative;
		border-radius: 4px;
		overflow: hidden;
	}

	.layout-logo {
		width: 100px;
		height: 30px;
		background: #5b6270;
		border-radius: 3px;
		float: left;
		position: relative;
		top: 15px;
		left: 20px;
	}

	.layout-nav {
		width: 420px;
		margin: 0 auto;
		margin-right: 20px;
	}

	.layout-footer-center {
		text-align: center;
	}
</style>
