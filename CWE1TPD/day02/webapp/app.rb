# app.rb
require 'sinatra'
require 'mysql2'
require 'json'
require 'yaml'

# Import files for database
client = Mysql2::Client.new(YAML.load_file('database.yml'))

get '/' do
  title_sql = "select * from films_title"
  @title_all = client.query(title_sql)
  erb :all, :layout => :mylayout
end

get '/new' do
  @new_hash = []
  for num in 1..10 do
    sql_new = "select * from films_title ORDER BY title_id desc limit 1 OFFSET #{num}" 
    sql_new = client.query(sql_new)
    @new_hash.push(sql_new)
  end
  erb :new, :layout => :mylayout
end

get '/post' do
  erb :post , :layout=> :mylayout
end

post '/confirm' do
  @post_title = params[:post_title]
  @post_impressions = params[:post_impressions]
  post_sql="INSERT INTO `films_title` (title,impressions) VALUES ('#{@post_title}','#{@post_impressions}');"
  client.query(post_sql)
  erb :confirm , :layout=> :mylayout
end

get '/search' do
  erb :search , :layout=> :mylayout
end

post '/search_re' do
  post_search = params[:post_search_re]
  search_sql="select * from films_title where title LIKE '%#{post_search}%'"
  #search_sql="select * from films_title where title LIKE '%ハ%'"
  @title_search = client.query(search_sql)
  erb :search_re , :layout=> :mylayout
end

get '/hello/*' do |name|
  "hello #{name} how are you?"
end

get '/title_id/*' do |id_name|
  title_sql = "select * from films_title where title_id = #{id_name}"
  @title_all = client.query(title_sql)
  erb :title_id, :layout => :mylayout
end

get '/category_id/*' do |id_category|
  title_sql = "select * from films_title where category_id = #{id_category}"
  @title_all = client.query(title_sql)
  erb :all, :layout => :mylayout
end

get "/all" do
  title_sql = "select * from films_title"
  @title_all = client.query(title_sql)
  erb :all, :layout => :mylayout
end

get '/saishin' do
  sql_saishin = "select * from films_title ORDER BY title_id desc limit 10"
  @title_aishin = client.query(sql_saishin)
  erb :saishin, :layout => :mylayout
end
