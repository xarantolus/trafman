create table Repositories (
    id int primary key,
    username text not null,
    name text not null,
    description text not null,
    is_fork boolean not null
);

create table RepoStats (
    repo_id int not null REFERENCES Repositories,
    date_time timestamp not null default CURRENT_TIMESTAMP,
    stars int not null,
    forks int not null,
    size int not null,
    watchers int not null,
    primary key (repo_id, date_time)
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
    date_time timestamp not null default CURRENT_TIMESTAMP,
    path text not null,
    title text not null,
    count int not null,
    uniques int not null,
    primary key (repo_id, date_time, path)
);

create table RepoTrafficReferrers (
    repo_id int not null REFERENCES Repositories,
    date_time timestamp not null default CURRENT_TIMESTAMP,
    referrer text not null,
    count int not null,
    uniques int not null,
    primary key (repo_id, date_time, referrer)
);

create table Releases (
    repo_id int not null REFERENCES Repositories,
    id int primary key not null,
    tag_name text not null,
    created timestamp not null default CURRENT_TIMESTAMP,
    name text not null,
    body text not null
);

create table ReleaseAssets (
    id int not null,
    release_id int not null REFERENCES Releases,
    primary key (id, release_id),
    date_time timestamp not null default CURRENT_TIMESTAMP,
    filename text not null,
    download_count int not null,
    updated_at date not null,
    size int not null
);
