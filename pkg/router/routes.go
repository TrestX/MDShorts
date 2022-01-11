package router

import (
	"MdShorts/pkg/services/account_service"
	"MdShorts/pkg/services/bookmark_service"
	"MdShorts/pkg/services/category_service"
	"MdShorts/pkg/services/news_service"
	"MdShorts/pkg/services/notification_service"
	"MdShorts/pkg/services/profile_service"
	"MdShorts/pkg/services/search_service"
	"MdShorts/pkg/services/share_service"
	"MdShorts/pkg/services/unregistered_user_service"
	user_news_check_service "MdShorts/pkg/services/userNewsCheck_service"
	"MdShorts/pkg/services/util_service"
	"net/http"

	"MdShorts/pkg/services/see_fewer_stories_service"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"signup",
		"POST",
		"/signup",
		account_service.SignUp,
	},
	Route{
		"login",
		"POST",
		"/login",
		account_service.Login,
	},
	Route{
		"loginsuccess",
		"POST",
		"/verify/emailotp",
		account_service.EmailVerifyOTP,
	},
	Route{
		"verifyemail",
		"GET",
		"/verifyemail/{code}",
		account_service.VerifyEmail,
	},
	Route{
		"verifymnumb",
		"POST",
		"/verifymobilenumber",
		account_service.VerifyOTP,
	},
	Route{
		"resendOtp",
		"POST",
		"/resendotp",
		account_service.ResendOTP,
	},
	Route{
		"sendemail",
		"POST",
		"/sendemail",
		account_service.SendVerificationEmail,
	},
	Route{
		"socialmedialogin",
		"POST",
		"/socialmedialogin",
		account_service.SocialMedialogin,
	},
	Route{
		"sendotp",
		"POST",
		"/resetpassword",
		account_service.SendPasswordResetOtp,
	},
	// Route{
	// 	"updatepassword",
	// 	"PUT",
	// 	"/updatepassword",
	// 	account_service.UpdatePassword,
	// },
	Route{
		"update profile",
		"PUT",
		"/profile",
		profile_service.UpdateProfile,
	},
	Route{
		"update profile",
		"PUT",
		"/password/profile",
		profile_service.ChangePassword,
	},
	Route{
		"set profile",
		"POST",
		"/profile",
		profile_service.SetProfile,
	},
	Route{
		"get profile",
		"GET",
		"/profile",
		profile_service.Profile,
	},
	Route{
		"getCategory",
		"GET",
		"/category",
		category_service.GetAllCategory,
	},
	Route{
		"updateCategory",
		"PUT",
		"/category/{categoryId}",
		category_service.UpdateCategory,
	},
	Route{
		"getCategoryByIds",
		"GET",
		"/category/{categoryIds}",
		category_service.GetCategoriesWithIDs,
	},
	Route{
		"category",
		"POST",
		"/category",
		category_service.AddCategory,
	},
	Route{
		"utilPreSigned",
		"POST",
		"/util/presignedurl",
		util_service.GetPreSignedURL,
	},
	Route{
		"getnews",
		"GET",
		"/news/{userId}",
		news_service.GetNews,
	},
	Route{
		"getGlobalnews",
		"GET",
		"/gnews",
		news_service.GetGlobalNews,
	},
	Route{
		"add news status for user",
		"POST",
		"/addnews",
		user_news_check_service.AddUserNewsCheck,
	},
	Route{
		"update news status for user",
		"PUT",
		"/updatenews",
		user_news_check_service.UpdateUserNewsCheck,
	},
	Route{
		"getnews",
		"GET",
		"/news/",
		news_service.GetNews,
	},
	Route{
		"getnewsbyID",
		"GET",
		"/newsbyID",
		news_service.GetNewsByID,
	},
	Route{
		"getnews",
		"GET",
		"/search/news",
		news_service.GetSearchNews,
	},
	Route{
		"getnews",
		"GET",
		"/topstories/news",
		news_service.GetTopStoriesNews,
	},
	Route{
		"getnews",
		"GET",
		"/trending/news",
		news_service.GetTrendingStoriesNews,
	},
	Route{
		"getnews",
		"GET",
		"/all/news",
		news_service.GetAllNews,
	},
	Route{
		"share",
		"POST",
		"/share",
		share_service.AddShare,
	},
	Route{
		"share",
		"GET",
		"/share",
		share_service.GetShares,
	},
	Route{
		"bookmark",
		"POST",
		"/bookmark",
		bookmark_service.AddBookmark,
	},
	Route{
		"bookmark",
		"GET",
		"/bookmark",
		bookmark_service.GetBookmarks,
	},
	Route{
		"bookmark",
		"PUT",
		"/bookmark/{bookmarkId}",
		bookmark_service.UpdateBookmark,
	},
	Route{
		"unregisteruser",
		"POST",
		"/add/unregisteruser",
		unregistered_user_service.AddUnregisteredUserService,
	},
	Route{
		"unregisteruser",
		"GET",
		"/unregisteruser",
		unregistered_user_service.AddUnregisteredUserService,
	},
	Route{
		"bookmarknewsids",
		"GET",
		"/bookmark/newsids",
		bookmark_service.GetBookmarksNewsIds,
	},
	Route{
		"unregisteruser",
		"GET",
		"/unregisteruser",
		unregistered_user_service.AddUnregisteredUserService,
	},
	Route{
		"search",
		"GET",
		"/search",
		search_service.GetSearch,
	},
	Route{
		"notify",
		"GET",
		"/notifications",
		notification_service.Getnotification,
	},
	Route{
		"notify",
		"POST",
		"/notifications",
		notification_service.SendNotificationWithTopic,
	},
	Route{
		"seefewer",
		"POST",
		"/seefewer",
		see_fewer_stories_service.AddSeeFewerStories,
	},
}
