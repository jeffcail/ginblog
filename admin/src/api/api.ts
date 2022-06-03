import axios from './axios';
//登录
export const login = (params: any) => {
    return axios.postJson(params,'/api/login')
};


//标签列表
export const getTagsList = () => {
    return axios.getJson("",'/api/tags/v1/get-tags-list')
};
// 创建标签
export const createTags = (params: any) => {
    return axios.postJson(params,'/api/tags/v1/create-tags')
};
// 修改标签
export const updateTagsById = (params: any) => {
    return axios.postJson(params,'/api/tags/v1/update-Tags-by-id')
};
// 删除标签
export const delTagsById = (params: any) => {
    return axios.postJson(params,'/api/tags/v1/del-tags-by-id')
};

// 分类列表
export const getCateList = (params: any) => {
    return axios.getJson(params,'/api/cate/v1/get-cate-list')
};
// 创建分类
export const createCate = (params: any) => {
    return axios.postJson(params,'/api/cate/v1/create-cate')
};
// 修改分类
export const updateCateById = (params: any) => {
    return axios.postJson(params,'/api/cate/v1/update-cate-by-id')
};
// 删除分类
export const delCateById = (params: any) => {
    return axios.postJson(params,'/api/cate/v1/del-cate-by-id')
};


// 获取文章列表
export const getPostsList = (params: any) => {
    return axios.getJson(params,'/api/posts/v1/get-posts-list')
};
// 文章创建
export const createPost = (params: any) => {
    return axios.postJson(params,'/api/posts/v1/create-post')
};
// 文章修改
export const updatePostById = (params: any) => {
    return axios.postJson(params,'/api/posts/v1/update-post-by-id')
};
// 文章删除
export const delPostById = (params: any) => {
    return axios.postJson(params,'/api/posts/v1/del-post-by-id')
};

// 获取评论列表
export const getCommentList = (params: any) => {
    return axios.getJson(params,'/api/comment/v1/get-comment-list')
};

// 获取友联列表
export const getLinkList = (params: any) => {
    return axios.getJson(params,'/api/link/v1/get-link-list')
};
// 友联创建
export const createLink = (params: any) => {
    return axios.postJson(params,'/api/link/v1/create-link')
};
// 友联修改
export const updateLinkById = (params: any) => {
    return axios.postJson(params,'/api/link/v1/update-link-by-id')
};
// 友联删除
export const delLinkById = (params: any) => {
    return axios.postJson(params,'/api/link/v1/del-link-by-id')
};







//账户管理
export const getUsersList = (params: any) => {
    return axios.getJson(params,'/api/user/v1/get-users')
};
// 查看
export const getUserById = (params: any) => {
    return axios.postJson(params,'/api/user/v1/get-user-by-id')
};
// 新增
export const createUsers = (params: any) => {
    return axios.postJson(params,'/api/user/v1/create-user')
};
// 删除
export const delUserById = (params: any) => {
    return axios.postJson(params,'/api/user/v1/del-user-by-id')
};
// 修改
export const updateUserById = (params: any) => {
    return axios.postJson(params,'/api/user/v1/update-user-by-id')
};
// 禁用
export const disableUserById = (params: any) => {
    return axios.postJson(params,'/api/user/v1/disable-user-by-id')
};
// 启用
export const enableUserById = (params: any) => {
    return axios.postJson(params,'/api/user/v1/enable-user-by-id')
};

