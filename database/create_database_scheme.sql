create table Repositories (
    id int primary key,
    username text not null,
    name text not null,
    description text not null,
    is_fork boolean not null
);

create table RepoStats (
    repo_id int not null REFERENCES Repositories,
    date date not null default CURRENT_DATE,
    stars int not null,
    forks int not null,
    size int not null,
    subscribers int not null,
    primary key (repo_id, date)
);

create table RepoTrafficViews (
    repo_id int not null REFERENCES Repositories,
    date date not null default CURRENT_DATE,
    count int not null,
    uniques int not null,
    primary key (repo_id, date)
);

create table RepoTrafficClones (
    repo_id int not null REFERENCES Repositories,
    date date not null default CURRENT_DATE,
    count int not null,
    uniques int not null,
    primary key (repo_id, date)
);

create table RepoTrafficPaths (
    repo_id int not null REFERENCES Repositories,
    date date not null default CURRENT_DATE,
    path text not null,
    title text not null,
    count int not null,
    uniques int not null,
    primary key (repo_id, date, path)
);

create table RepoTrafficReferrers (
    repo_id int not null REFERENCES Repositories,
    date date not null default CURRENT_DATE,
    referrer text not null,
    count int not null,
    uniques int not null,
    primary key (repo_id, date, referrer)
);


create table Releases (
    repo_id int not null REFERENCES Repositories,

    id int primary key not null,
    tag_name text not null,
    created date not null default CURRENT_DATE,

    name text not null,
    body text not null
);

create table ReleaseAssets (
    id int not null,
    release_id int not null REFERENCES Releases,
    primary key (id, release_id),

    date date not null default CURRENT_DATE,

    filename text not null,
    download_count int not null,
    updated_at date not null,
    size int not null
);
