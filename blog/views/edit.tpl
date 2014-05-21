 <h1>编辑 {{.Post.Title}}</h1>
 <form action="" method="post">
 标题:<input type="text" name="title" class="form-control" value="{{.Post.Title}}">
 内容：<textarea name="content" class="xheditor form-control" rows="10">{{.Post.Content}}</textarea>
 <input type="hidden" name="id" class="form-control"  value="{{.Post.Id}}">
<input type="submit" value="保存" />
 </form>
