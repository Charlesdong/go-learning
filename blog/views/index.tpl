<div class="list-group">
{{range .blogs}}
<div class="col-md-10">
  <a href="/view/{{.Id}}" class="list-group-item">
    <h4 class="list-group-item-heading">{{.Title}}</h4>
    <p class="list-group-item-text">{{cutstring (html2str .Content) 30 }}...</p>
    <p class="list-group-item-text">{{date .Created "Y-m-d H:i:s"}}</p>
  </a>
</div>
<div class="col-md-2">
<ul class="list-unstyled">
<li></li>
<li> <a href="/edit/{{.Id}}" class="list-group-item">编辑</a> </li> 
<li> <a href="/delete/{{.Id}}" class="list-group-item">删除</a> </li>
</ul>
</div>
{{end}}
</div>
