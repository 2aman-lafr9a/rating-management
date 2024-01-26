# **Rating Management:**

These 2 services were written in GOLANG and used a PostgreSQL database, the Rating micro-service is responsible for managing Rating. It provides functionalities such as creating, updating and retrieving agencies.

[https://en.wikipedia.org/wiki/Go_(programming_language)](https://en.wikipedia.org/wiki/Go_(programming_language))

It uses a PostgreSQL database ( Relational database ) as  a primary database to store its data, and it provides a GRPC interface to deal with it, it’s dumps and all that it knows is Rating only, the methods it provides are
## Rating
	 rpc FindById(GetRatingRequest) returns (GetRatingResponse);
    rpc FindAll(GetRatingsRequest) returns (GetRatingsResponse);
    rpc Update(UpdateRatingRequest) returns (UpdateRatingResponse);
