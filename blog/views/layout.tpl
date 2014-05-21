<html>
<head>
<title>My Blog</title>
<style>
#menu {
width: 200px;
float: right;
}
</style>
<link rel="stylesheet" href="/static/bootstrap/3.0.3/css/bootstrap.min.css">
<!-- 可选的Bootstrap主题文件（一般不用引入） -->
<link rel="stylesheet" href="/static/bootstrap/3.0.3/css/bootstrap-theme.min.css">
</head>

<body>
<div class="container">
  <!-- Static navbar -->
      <div class="navbar navbar-default" role="navigation">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="/">TATA'S Blog</a>
        </div>
        <div class="navbar-collapse collapse">
          <ul class="nav navbar-nav">
            <li class="active"><a href="/">Home</a></li>
            <li><a href="/new">New</a></li>
          </ul>
        </div><!--/.nav-collapse -->
      </div>
<div class="row">
{{.LayoutContent}}
</div>
      <hr>
      <footer>
        <p>Tata&copy; Company 2013</p>
      </footer>
</div>

<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
<!--script src="http://cdn.bootcss.com/jquery/1.10.2/jquery.min.js"></script -->
<script type="text/javascript" src="/static/js/jquery-1.4.4.min.js"></script>

<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
<script src="/static/bootstrap/3.0.3/js/bootstrap.min.js"></script>
<script type="text/javascript" src="/static/js/xheditor/xheditor-1.2.1.min.js"></script>
<script type="text/javascript" src="/static/js/xheditor/xheditor_lang/zh-cn.js"></script>
</body>
</html>

